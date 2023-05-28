package main

import (
	"bufio"
	"lab4/config"
	"lab4/test"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func apiTest(url string, courseList []test.CourseItem, studentList []test.StudentItem, mode string) {
	log.Printf("[Test] -----%s测试开始-----\n", mode)
	log.Printf("[Test]\x1b[35m Test /api/search/course Start\x1b[0m\n")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(courseList) - 1)
	courseId := courseList[num].ID
	log.Printf("[Test] 查询课程ID: %v\n", courseId)
	course, err := test.GetCourse(url+"/api/search/course", courseId)
	if err != nil {
		log.Printf("[Test] 查询课程测试失败: %v\n", err)
	} else {
		if course.ID != "" {
			log.Printf("[Test] 查询课程测试成功\n")
			log.Printf("[Test] 课程编号: %v\n", course.ID)
			log.Printf("[Test] 课程名称: %v\n", course.Name)
			log.Printf("[Test] 课程容量: %v\n", course.Capacity)
			log.Printf("[Test] 已选人数: %v\n", course.Selected)
			log.Printf("[Test]\x1b[32m /api/search/course PASS\x1b[0m\n")
		} else {
			log.Printf("[Test] 查询课程测试失败: 未查到有效课程信息\n")
		}
	}
	log.Printf("[Test]\x1b[35m Test /api/search/course End\x1b[0m\n\n")

	log.Printf("[Test]\x1b[35m Test /api/search/all Start\x1b[0m\n")
	allCourse, err := test.GetAllCourse(url + "/api/search/all")
	if err != nil {
		log.Printf("[Test] 查询全部课程测试失败: %v\n", err)
	} else {
		if len(*allCourse) > 0 {
			log.Print("[Test] 查询全部课程测试成功\n")
			log.Printf("[Test] 课程总数: %d门\n", len(*allCourse))
			log.Printf("[Test]\x1b[32m /api/search/all PASS\x1b[0m\n")
		} else {
			log.Printf("[Test] 查询全部课程测试失败: 所有课程为空\n")
		}
	}
	log.Printf("[Test]\x1b[35m Test /api/search/all End\x1b[0m\n\n")

	log.Printf("[Test]\x1b[35m Test /api/search/student Start\x1b[0m\n")
	num = r.Intn(len(studentList) - 1)
	stuid := studentList[num].Stuid
	log.Printf("[Test] 查询学生ID: %v\n", stuid)
	studentCourse, err := test.GetStudentCourse(url+"/api/search/student", stuid)
	if err != nil {
		log.Printf("[Test] 查询学生课程测试失败: %v\n", err)
	} else {
		if studentCourse.Stuid == stuid {
			log.Print("[Test] 查询学生课程测试成功")
			log.Printf("[Test] 学号：%v\n", studentCourse.Stuid)
			log.Printf("[Test] 姓名：%v\n", studentCourse.Name)
			log.Printf("[Test] 课程: %v\n", studentCourse.Course)
			log.Printf("[Test]\x1b[32m /api/search/student PASS\x1b[0m\n")
		} else {
			log.Printf("[Test] 查询学生课程测试失败: 未查到有效学生信息\n")
		}
	}
	log.Printf("[Test]\x1b[35m Test /api/search/student End\x1b[0m\n\n")

	log.Printf("[Test]\x1b[35m Test /api/choose Start\x1b[0m\n")
	err = test.ChooseCourse(url+"/api/choose", stuid, courseId)
	if err != nil {
		log.Printf("[Test] 选课测试失败: %v\n", err)
	} else {
		studentCourse, err := test.GetStudentCourse(url+"/api/search/student", stuid)
		if err != nil {
			log.Printf("[Test] 选课测试失败: %v\n", err)
		} else {
			if len(studentCourse.Course) > 0 && studentCourse.Course[0].ID == courseId {
				log.Printf("[Test] 选课测试成功: %v\n", course)
				log.Printf("[Test] 学生课程信息: %v\n", studentCourse.Course)
				log.Printf("[Test]\x1b[32m /api/choose PASS\x1b[0m\n")
			} else {
				log.Printf("[Test] 学生课程信息: %v\n", studentCourse.Course)
				log.Printf("[Test] 选课测试失败: 未查到有效选课结果\n")
			}
		}
	}
	log.Printf("[Test]\x1b[35m Test /api/choose End\x1b[0m\n\n")

	log.Printf("[Test]\x1b[35m Test /api/drop Start\x1b[0m\n")
	err = test.DropCourse(url+"/api/drop", stuid, courseId)
	if err != nil {
		log.Printf("[Test] 退课测试失败: %v\n", err)
	} else {
		studentCourse, err := test.GetStudentCourse(url+"/api/search/course", stuid)
		if err != nil {
			log.Printf("[Test] 退课测试失败: %v\n", err)
		} else {
			if len(studentCourse.Course) == 0 {
				log.Printf("[Test] 退课测试成功\n")
				log.Printf("[Test] 学生课程信息: %v\n", studentCourse.Course)
				log.Printf("[Test]\x1b[32m /api/drop PASS\x1b[0m\n")
			} else {
				log.Printf("[Test] 学生课程信息: %v\n", studentCourse.Course)
				log.Printf("[Test] 退课测试失败: 课程未退选\n")
			}
		}
	}
	log.Printf("[Test]\x1b[35m Test /api/drop End\x1b[0m\n\n")
	log.Printf("[Test] -----%s测试结束-----\n\n", mode)
}

func startProcess(rootPath string, cmdStr string) *exec.Cmd {
	cmd := exec.Command("/bin/bash", "-c", "cd "+rootPath+" && "+cmdStr+"")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, //使得 Shell 进程开辟新的 PGID, 即 Shell 进程的 PID, 它后面创建的所有子进程都属于该进程组
	}
	if config.GetStatic().Debug {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}
	if err := cmd.Start(); err != nil {
		log.Fatalf("[Test] %s cmd.Start error: %v", cmdStr, err)
	}

	log.Printf("[Test]\x1b[32m Run Procsee is %v\x1b[0m\n", cmdStr)
	log.Printf("[Test]\x1b[32m Run Success, Pid is %d\x1b[0m\n", cmd.Process.Pid)
	return cmd
}
func endProcess(cmd *exec.Cmd) {
	err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
	if err != nil {
		log.Printf("[Test]\x1b[31m Process %d exit failed: %v\x1b[0m\n\n", cmd.Process.Pid, err)
	}

}

func build(c config.StaticConfig) {
	log.Printf("[Test]\x1b[35m Build Start\x1b[0m\n\n")
	cmd := exec.Command("/bin/bash", "-c", "cd "+c.RootPath+" && "+c.Clean+" && "+c.Build)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Fatalf("[Test]\x1b[31m build error: %v\x1b[0m", err)
	}
	log.Printf("[Test]\x1b[32m Build Success\x1b[0m\n")

}

func readStudent(filePath string) []test.StudentItem {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("[Test]\x1b[31m student file read error: %v\x1b[0m", err)
	}
	var list []test.StudentItem
	buf := bufio.NewScanner(f)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line = strings.TrimSpace(line)
		strSlice := strings.Split(line, " ")
		var tmp test.StudentItem
		tmp.Stuid = strSlice[0]
		tmp.Name = strSlice[1]
		list = append(list, tmp)
	}
	return list
}

func readCourse(filePath string) []test.CourseItem {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("[Test]\x1b[31m course file read error: %v\x1b[0m", err)
	}
	var list []test.CourseItem
	buf := bufio.NewScanner(f)
	for {
		if !buf.Scan() {
			break
		}
		line := buf.Text()
		line = strings.TrimSpace(line)
		strSlice := strings.Split(line, " ")
		var tmp test.CourseItem
		tmp.ID = strSlice[0]
		tmp.Name = strSlice[1]
		tmp.Capacity, _ = strconv.Atoi(strSlice[2])
		tmp.Selected = 0
		list = append(list, tmp)
	}
	return list
}

func main() {
	config.Init()
	config.Check()

	// 编译可执行文件
	build(config.GetStatic())

	// 启动Web / 负载均衡
	var webCmd *exec.Cmd
	var balanceCmd *exec.Cmd
	var balanceWebCmd []*exec.Cmd
	webNum := len(config.GetBalancer().CMD)
	url := ""
	if config.GetBalancer().Status {
		balanceCmd = startProcess(config.GetStatic().RootPath, config.GetBalancer().CMD)
		url = config.GetBalancer().URL
		balanceWebCmd = make([]*exec.Cmd, webNum)
		for i := 0; i < webNum; i++ {
			balanceWebCmd[i] = startProcess(config.GetStatic().RootPath, config.GetStore().TPC.CMD[i])
		}
	} else {
		webCmd = startProcess(config.GetStatic().RootPath, config.GetWeb().CMD)
		url = config.GetWeb().URL
	}

	// 启动存储节点
	storeNum := 0
	var storeCmd []*exec.Cmd
	if config.GetStore().TPC.Status { // 2PC模式
		storeNum = len(config.GetStore().TPC.CMD)
		storeCmd = make([]*exec.Cmd, storeNum)
		for i := 0; i < storeNum; i++ {
			storeCmd[i] = startProcess(config.GetStatic().RootPath, config.GetStore().TPC.CMD[i])
		}
	} else { // RAFT模式
		storeNum = len(config.GetStore().RAFT.CMD)
		storeCmd = make([]*exec.Cmd, storeNum)
		for i := 0; i < storeNum; i++ {
			storeCmd[i] = startProcess(config.GetStatic().RootPath, config.GetStore().RAFT.CMD[i])
		}
	}
	// 读取课程数据
	courseList := readCourse(config.GetStatic().RootPath + "/" + config.GetStatic().Course)

	// 读取学生数据
	studentList := readStudent(config.GetStatic().RootPath + "/" + config.GetStatic().Student)

	// 开始basic测试
	apiTest(url, courseList, studentList, "API")

	// 销毁存储节点
	endProcess(storeCmd[2])
	endProcess(storeCmd[6])

	// 测试API功能
	apiTest(url, courseList, studentList, "Store")

	if config.GetBalancer().Status {
		// 销毁Web节点
		endProcess(balanceWebCmd[0])

		// 测试API功能
		apiTest(url, courseList, studentList, "Balancer")
	}

	// 销毁所有进程
	if config.GetBalancer().Status {
		endProcess(balanceCmd)
		for i := 0; i < webNum; i++ {
			endProcess(balanceWebCmd[i])
		}
	} else {
		endProcess(webCmd)
	}

	// 销毁存储节点进程
	for i := 0; i < storeNum; i++ {
		endProcess(storeCmd[i])
	}
}
