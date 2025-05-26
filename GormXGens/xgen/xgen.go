package xgen

import (
	"common-toolkits-v1/GormXGens/config"
	"common-toolkits-v1/logx/logs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

// XGen XGen对象
type XGen struct {
	xGenConfig config.XGenConfig // 配置文件
	db         *gorm.DB          // gorm连接
}

// NewXGen 已经初始化了
func NewXGen(XGenConfig config.XGenConfig) *XGen {
	dsn := XGenConfig.DNS
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logs.Error(err)
	}
	return &XGen{
		xGenConfig: XGenConfig,
		db:         db,
	}
}

// InitDB 初始化DB
func (x *XGen) InitDB() error {
	dsn := x.xGenConfig.DNS
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logs.Error(err)
	}
	x.db = db
	return nil
}

// CreateXGenModel 生成Model结构
func (x *XGen) CreateXGenModel() error {
	// 初始化gorm的Gen
	g := gen.NewGenerator(x.xGenConfig.GenConfig)
	g.UseDB(x.db)

	for _, value := range x.xGenConfig.GenModelList {
		g.GenerateModel(value)
	}
	g.Execute()

	return nil
}

// CreateXGenQuery 生成模板操作
func (x *XGen) CreateXGenQuery() error {

	// 获取需要更新model文件的位置
	modelListPaths, err := x.getGenModelListPaths()
	if err != nil {
		return err
	}

	imp, err := x.getXGenQueryImportModelPath()
	if err != nil {
		return err
	}

	for k, p := range modelListPaths {
		modelObject, err := x.reflectXGenModel(p)
		if err != nil {
			return err
		}

		doName := fmt.Sprintf("%vQuery", modelObject.ModelName)

		// 处理ModelPackage
		cleanedPath := filepath.Clean(x.xGenConfig.GenConfig.ModelPkgPath)
		cleanedPath = strings.TrimSuffix(cleanedPath, "/")
		modelPackage := filepath.Base(cleanedPath)

		ParentName := lowercaseFirst(doName)
		ExtendName := doName

		PrimaryKey, IndexList, err := x.getKeyData(*modelObject)
		if err != nil {
			logs.Error("get key data err:", err)
			return err
		}

		fileTemp := config.TemplateField{
			QueryName:     ParentName,
			QueryPackage:  x.xGenConfig.QueryPkgName,
			ModelName:     modelObject.ModelName,
			ModelPackage:  modelPackage,
			ModelLinkPath: imp,
			PrimaryKey:    PrimaryKey,
			IndexList:     IndexList,
		}

		queryFile := fmt.Sprintf("%v_query.gen.go", x.xGenConfig.GenModelList[k])
		createPath := path.Join(x.xGenConfig.QueryPkgPath, x.xGenConfig.GenModelList[k], queryFile)

		err = x.createXGenTemplate(fileTemp, "../template/xgen_template.tpl", createPath)
		if err != nil {
			return err
		}

		fieldExtend := config.TemplateFieldExtend{
			ParentName:   ParentName,
			ExtendName:   ExtendName,
			QueryPackage: x.xGenConfig.QueryPkgName,
		}
		extendFile := fmt.Sprintf("%v_extend.go", x.xGenConfig.GenModelList[k])
		extendPath := path.Join(x.xGenConfig.QueryPkgPath, x.xGenConfig.GenModelList[k], extendFile)

		err = x.createXGenTemplate(fieldExtend, "../template/xgen_template_extend.tpl", extendPath)
		if err != nil {
			return err
		}
	}

	return nil
}

// 根据filed和模板生成对应操作文件
func (x *XGen) createXGenTemplate(filed any, templatePath, createPath string) error {

	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		logs.Error(err)
		return err
	}

	parseTemp, err := template.New("xgen_template").Parse(string(templateContent))
	if err != nil {
		logs.Errorf("解析模板出错: %v\n", err)
		return err
	}

	// 提取目录部分
	dir := filepath.Dir(createPath)

	// 创建目录（包括所有必要的父目录）
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("创建目录失败: %v\n", err)
		return err
	}

	// 执行模板并将结果输出到文件
	file, err := os.Create(createPath)
	if err != nil {
		logs.Errorf("创建文件出错: %v\n", err)
		return err
	}
	defer file.Close()

	err = parseTemp.Execute(file, filed)
	if err != nil {
		logs.Errorf("执行模板出错: %v\n", err)
		return err
	}
	return nil
}

// 反射Model结构生成代码
func (x *XGen) reflectXGenModel(modelFilePath string) (*ModelObject, error) {

	// 解析文件
	modelObject, err := ParseStructFromFile(modelFilePath)
	if err != nil {
		return nil, err
	}

	return modelObject, nil
}

// 生成Query操作关联的Model的位置链接
func (x *XGen) getXGenQueryImportModelPath() (string, error) {

	projectName := x.xGenConfig.ProjectPath

	currentPath, err := getCurrentLocationPath(projectName)
	if err != nil {
		return "", err
	}

	goModPath, err := getGoModName()
	if err != nil {
		return "", err
	}
	pathData := path.Join(goModPath, currentPath)

	modelPath := processRelativePath(pathData, x.xGenConfig.ModelPkgPath)

	return modelPath, nil
}

// 获取路径下的全部文件名称
func (x *XGen) getModelFilePaths() ([]string, error) {
	modelDirPath := x.xGenConfig.ModelPkgPath

	// 检查目录是否存在
	if _, err := os.Stat(modelDirPath); os.IsNotExist(err) {
		logs.Errorf("错误：目录 %s 不存在\n", modelDirPath)
		return nil, err
	} else if err != nil {
		logs.Errorf("错误：无法访问目录 %s: %v\n", modelDirPath, err)
		return nil, err
	}

	modelPathList := make([]string, 0)

	// 遍历目录树
	err := filepath.WalkDir(modelDirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件（非目录），则打印路径
		if !d.IsDir() {
			modelPathList = append(modelPathList, path)
			fmt.Println(path)
		}
		return nil
	})

	if err != nil {
		logs.Errorf("遍历目录时出错: %v\n", err)
		return nil, err
	}

	return modelPathList, nil

}

// 通过更新表单获取跟新路径列表
func (x *XGen) getGenModelListPaths() ([]string, error) {
	modelList := x.xGenConfig.GenModelList
	modelDirPath := x.xGenConfig.ModelPkgPath

	modelPathList := make([]string, 0)
	for _, modelName := range modelList {
		modelName = fmt.Sprintf("%v.gen.go", modelName)
		p := path.Join(modelDirPath, modelName)
		modelPathList = append(modelPathList, p)
	}
	return modelPathList, nil
}

func (x *XGen) getKeyData(modelObject ModelObject) (*config.PrimaryKeyData, []config.IndexKeyData, error) {

	IndexList := make([]config.IndexKeyData, 0)
	isPrimaryKey := false
	var primaryKeyData *config.PrimaryKeyData

	for _, filed := range modelObject.FieldMap {

		// 主键处理
		if filed.GormTag.IsPrimaryKey {
			primaryKeyData = &config.PrimaryKeyData{
				Name:   filed.FieldName,
				Type:   fmt.Sprint(filed.FieldType),
				Column: filed.FieldColumn,
			}
			isPrimaryKey = true
			continue
		}

		// 索引处理
		if filed.GormTag.IsIndex {
			data := config.IndexKeyData{
				Name:   filed.FieldName,
				Type:   fmt.Sprint(filed.FieldType),
				Column: filed.FieldColumn,
			}
			IndexList = append(IndexList, data)
		}

	}

	if !isPrimaryKey {
		return nil, IndexList, nil
	}

	return primaryKeyData, IndexList, nil
}
