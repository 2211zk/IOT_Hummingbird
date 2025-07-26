package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// TestRunner 测试运行器
type TestRunner struct {
	rootDir string
}

// NewTestRunner 创建测试运行器
func NewTestRunner(rootDir string) *TestRunner {
	return &TestRunner{rootDir: rootDir}
}

// RunAllTests 运行所有测试
func (tr *TestRunner) RunAllTests() error {
	fmt.Println("🚀 开始运行部门管理模块的所有测试...")

	// 定义测试套件
	testSuites := []TestSuite{
		{
			Name:        "Service层单元测试",
			Path:        "service/wl_department",
			TestFile:    "wl_department_test.go",
			Description: "测试Service层业务逻辑",
		},
		{
			Name:        "API层单元测试",
			Path:        "api/v1/wl_department",
			TestFile:    "wl_department_test.go",
			Description: "测试API接口层",
		},
		{
			Name:        "树形结构单元测试",
			Path:        "test/unit",
			TestFile:    "wl_department_tree_test.go",
			Description: "测试部门树形结构和循环引用检测",
		},
		{
			Name:        "设备关联单元测试",
			Path:        "test/unit",
			TestFile:    "wl_department_device_test.go",
			Description: "测试设备关联功能",
		},
		{
			Name:        "集成测试",
			Path:        "test/integration",
			TestFile:    "wl_department_integration_test.go",
			Description: "测试完整业务流程",
		},
	}

	var failedTests []string
	var passedTests []string

	for _, suite := range testSuites {
		fmt.Printf("\n📋 运行测试套件: %s\n", suite.Name)
		fmt.Printf("   描述: %s\n", suite.Description)
		fmt.Printf("   路径: %s\n", suite.Path)

		if err := tr.runTestSuite(suite); err != nil {
			fmt.Printf("   ❌ 失败: %v\n", err)
			failedTests = append(failedTests, suite.Name)
		} else {
			fmt.Printf("   ✅ 通过\n")
			passedTests = append(passedTests, suite.Name)
		}
	}

	// 输出测试总结
	tr.printTestSummary(passedTests, failedTests)

	if len(failedTests) > 0 {
		return fmt.Errorf("有 %d 个测试套件失败", len(failedTests))
	}

	return nil
}

// RunSpecificTest 运行特定测试
func (tr *TestRunner) RunSpecificTest(testName string) error {
	fmt.Printf("🎯 运行特定测试: %s\n", testName)

	testSuites := tr.getTestSuites()

	for _, suite := range testSuites {
		if strings.Contains(strings.ToLower(suite.Name), strings.ToLower(testName)) {
			return tr.runTestSuite(suite)
		}
	}

	return fmt.Errorf("未找到匹配的测试套件: %s", testName)
}

// RunTestWithCoverage 运行测试并生成覆盖率报告
func (tr *TestRunner) RunTestWithCoverage() error {
	fmt.Println("📊 运行测试并生成覆盖率报告...")

	// 创建覆盖率输出目录
	coverageDir := filepath.Join(tr.rootDir, "test", "coverage")
	if err := os.MkdirAll(coverageDir, 0755); err != nil {
		return fmt.Errorf("创建覆盖率目录失败: %v", err)
	}

	// 运行覆盖率测试
	coverageFile := filepath.Join(coverageDir, "coverage.out")
	cmd := exec.Command("go", "test", "-coverprofile="+coverageFile, "./...")
	cmd.Dir = tr.rootDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("覆盖率测试输出:\n%s\n", string(output))
		return fmt.Errorf("覆盖率测试失败: %v", err)
	}

	fmt.Printf("覆盖率测试输出:\n%s\n", string(output))

	// 生成HTML覆盖率报告
	htmlFile := filepath.Join(coverageDir, "coverage.html")
	cmd = exec.Command("go", "tool", "cover", "-html="+coverageFile, "-o", htmlFile)
	cmd.Dir = tr.rootDir

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("生成HTML覆盖率报告失败: %v", err)
	}

	fmt.Printf("✅ 覆盖率报告已生成: %s\n", htmlFile)
	return nil
}

// RunBenchmarks 运行性能测试
func (tr *TestRunner) RunBenchmarks() error {
	fmt.Println("⚡ 运行性能测试...")

	benchmarkPaths := []string{
		"./service/wl_department",
		"./api/v1/wl_department",
	}

	for _, path := range benchmarkPaths {
		fmt.Printf("\n📈 运行 %s 的性能测试...\n", path)

		cmd := exec.Command("go", "test", "-bench=.", "-benchmem", path)
		cmd.Dir = tr.rootDir

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("性能测试失败: %v\n输出: %s\n", err, string(output))
			continue
		}

		fmt.Printf("性能测试结果:\n%s\n", string(output))
	}

	return nil
}

// runTestSuite 运行单个测试套件
func (tr *TestRunner) runTestSuite(suite TestSuite) error {
	testPath := filepath.Join(tr.rootDir, suite.Path)

	// 检查测试文件是否存在
	testFile := filepath.Join(testPath, suite.TestFile)
	if _, err := os.Stat(testFile); os.IsNotExist(err) {
		return fmt.Errorf("测试文件不存在: %s", testFile)
	}

	// 运行测试
	cmd := exec.Command("go", "test", "-v", ".")
	cmd.Dir = testPath

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("测试输出:\n%s\n", string(output))
		return fmt.Errorf("测试执行失败: %v", err)
	}

	// 解析测试结果
	tr.parseTestOutput(string(output))

	return nil
}

// parseTestOutput 解析测试输出
func (tr *TestRunner) parseTestOutput(output string) {
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.Contains(line, "PASS:") {
			fmt.Printf("     ✅ %s\n", line)
		} else if strings.Contains(line, "FAIL:") {
			fmt.Printf("     ❌ %s\n", line)
		} else if strings.Contains(line, "RUN") {
			fmt.Printf("     🏃 %s\n", line)
		} else if strings.Contains(line, "ok") && strings.Contains(line, "coverage:") {
			fmt.Printf("     📊 %s\n", line)
		}
	}
}

// printTestSummary 打印测试总结
func (tr *TestRunner) printTestSummary(passed, failed []string) {
	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("📋 测试总结")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Printf("✅ 通过的测试套件 (%d):\n", len(passed))
	for _, test := range passed {
		fmt.Printf("   • %s\n", test)
	}

	if len(failed) > 0 {
		fmt.Printf("\n❌ 失败的测试套件 (%d):\n", len(failed))
		for _, test := range failed {
			fmt.Printf("   • %s\n", test)
		}
	}

	fmt.Printf("\n📊 总计: %d 通过, %d 失败\n", len(passed), len(failed))

	if len(failed) == 0 {
		fmt.Println("🎉 所有测试都通过了！")
	} else {
		fmt.Println("⚠️  有测试失败，请检查上面的错误信息")
	}
}

// getTestSuites 获取所有测试套件
func (tr *TestRunner) getTestSuites() []TestSuite {
	return []TestSuite{
		{
			Name:        "Service层单元测试",
			Path:        "service/wl_department",
			TestFile:    "wl_department_test.go",
			Description: "测试Service层业务逻辑",
		},
		{
			Name:        "API层单元测试",
			Path:        "api/v1/wl_department",
			TestFile:    "wl_department_test.go",
			Description: "测试API接口层",
		},
		{
			Name:        "树形结构单元测试",
			Path:        "test/unit",
			TestFile:    "wl_department_tree_test.go",
			Description: "测试部门树形结构和循环引用检测",
		},
		{
			Name:        "设备关联单元测试",
			Path:        "test/unit",
			TestFile:    "wl_department_device_test.go",
			Description: "测试设备关联功能",
		},
		{
			Name:        "集成测试",
			Path:        "test/integration",
			TestFile:    "wl_department_integration_test.go",
			Description: "测试完整业务流程",
		},
	}
}

// TestSuite 测试套件定义
type TestSuite struct {
	Name        string
	Path        string
	TestFile    string
	Description string
}

// main 主函数
func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// 获取项目根目录
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取当前目录失败: %v\n", err)
		os.Exit(1)
	}

	// 确保在server目录下运行
	if !strings.HasSuffix(rootDir, "server") {
		rootDir = filepath.Join(rootDir, "server")
	}

	runner := NewTestRunner(rootDir)

	command := os.Args[1]

	switch command {
	case "all":
		if err := runner.RunAllTests(); err != nil {
			fmt.Printf("运行所有测试失败: %v\n", err)
			os.Exit(1)
		}
	case "coverage":
		if err := runner.RunTestWithCoverage(); err != nil {
			fmt.Printf("运行覆盖率测试失败: %v\n", err)
			os.Exit(1)
		}
	case "bench":
		if err := runner.RunBenchmarks(); err != nil {
			fmt.Printf("运行性能测试失败: %v\n", err)
			os.Exit(1)
		}
	case "test":
		if len(os.Args) < 3 {
			fmt.Println("请指定要运行的测试名称")
			os.Exit(1)
		}
		testName := os.Args[2]
		if err := runner.RunSpecificTest(testName); err != nil {
			fmt.Printf("运行特定测试失败: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("未知命令: %s\n", command)
		printUsage()
		os.Exit(1)
	}

	fmt.Println("\n🎯 测试运行完成！")
}

// printUsage 打印使用说明
func printUsage() {
	fmt.Println("部门管理模块测试运行器")
	fmt.Println("")
	fmt.Println("用法:")
	fmt.Println("  go run test/run_tests.go <command> [args]")
	fmt.Println("")
	fmt.Println("命令:")
	fmt.Println("  all              运行所有测试")
	fmt.Println("  coverage         运行测试并生成覆盖率报告")
	fmt.Println("  bench            运行性能测试")
	fmt.Println("  test <name>      运行特定的测试套件")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  go run test/run_tests.go all")
	fmt.Println("  go run test/run_tests.go coverage")
	fmt.Println("  go run test/run_tests.go bench")
	fmt.Println("  go run test/run_tests.go test service")
	fmt.Println("  go run test/run_tests.go test api")
}
