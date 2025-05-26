package xgen

import (
	"common-toolkits-v1/logx/logs"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getCurrentLocationPath(ProjectName string) (string, error) {
	pathData, err := os.Getwd()
	if err != nil {
		logs.Error("路径获取错误：", err)
		return "", nil
	}

	// 转换路径
	cleanPath := filepath.Clean(pathData)

	components := strings.Split(cleanPath, string(filepath.Separator))

	// 查找项目名称在路径中的位置
	foundIndex := -1
	for i, component := range components {
		if component == ProjectName {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		errStr := fmt.Sprintf("未找到项目目录 %v", ProjectName)
		logs.Error(errStr)
		return "", errors.New(errStr)
	}

	// 提取从项目目录开始的路径组件
	resultComponents := components[foundIndex+1:]

	// 重新组合路径
	result := strings.Join(resultComponents, string(filepath.Separator))

	// 转换为 Unix 风格路径（可选）
	result = strings.ReplaceAll(result, string(filepath.Separator), "/")

	return result, nil
}

func getGoModName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	output, err := cmd.Output()
	if err != nil {
		logs.Errorf("go.mod 名称获取错误：%v", err)
		return "", err
	}

	goModPath := strings.TrimSpace(string(output))
	return goModPath, nil
}

func processRelativePath(basePath, inputPath string) string {

	// 统计开头的 ../ 数量
	var dotDotCount int
	for strings.HasPrefix(inputPath, "../") {
		dotDotCount++
		inputPath = strings.TrimPrefix(inputPath, "../")
	}

	// 标准化路径分隔符
	basePath = filepath.Clean(basePath)
	inputPath = filepath.Clean(inputPath)

	// 如果没有 ../，直接返回输入路径
	if dotDotCount == 0 {
		return inputPath
	}

	// 将基准路径分割为目录列表
	baseDirs := strings.Split(basePath, string(filepath.Separator))

	// 确保基准路径有足够的目录可退
	if dotDotCount > len(baseDirs) {
		dotDotCount = len(baseDirs)
	}

	// 移除基准路径末尾的 N 个目录
	adjustedBaseDirs := baseDirs[:len(baseDirs)-dotDotCount]

	// 拼接调整后的基准路径和剩余的输入路径
	result := filepath.Join(append(adjustedBaseDirs, inputPath)...)

	result = filepath.Clean(result)

	result = strings.ReplaceAll(result, string(filepath.Separator), "/")
	return result
}
