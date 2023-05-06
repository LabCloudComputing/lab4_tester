package test

import (
	"errors"
	"lab4/config"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
)

func showTrace(resp *resty.Response) {
	if config.GetStatic().HttpDebug {
		log.Printf("[HTTP] |\x1b[42;37m %d \x1b[0m| %v |\x1b[43;37m %v \x1b[0m| %v | %v\n", resp.StatusCode(), resp.Request.URL, resp.Request.Method, resp.Request.Body, resp)
	}
}
func GetCourse(url string, course_id string) (*CourseItem, error) {
	client := resty.New()
	var result Course
	client = client.SetRetryCount(1).SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second)
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"id": course_id,
		}).
		SetResult(&result).
		ForceContentType("application/json").
		EnableTrace().
		Get(url)
	showTrace(resp)
	if err != nil {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, err
	}
	if result.Code != 200 {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, errors.New("响应码错误 - " + result.Msg)
	}
	return &result.Data, nil
}

func GetAllCourse(url string) (*[]CourseItem, error) {
	client := resty.New()
	var result AllCourse
	resp, err := client.R().
		SetResult(&result).
		ForceContentType("application/json").
		EnableTrace().
		Get(url)
	showTrace(resp)
	if err != nil {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, err
	}
	if result.Code != 200 {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, errors.New("响应码错误 - " + result.Msg)
	}
	return &result.Data, nil
}

func GetStudentCourse(url string, stuid string) (*StudentCourseItem, error) {
	client := resty.New()
	var result StudentCourse
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"stuid": stuid,
		}).
		SetResult(&result).
		ForceContentType("application/json").
		EnableTrace().
		Get(url)
	showTrace(resp)
	if err != nil {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, err
	}
	if result.Code != 200 {
		log.Printf("[HTTP] Response: %v\n", resp)
		return nil, errors.New("响应码错误 - " + result.Msg)
	}
	return &result.Data, nil
}

func ChooseCourse(url string, stuid string, course_id string) error {
	client := resty.New()
	var result CourseResult
	resp, err := client.R().
		SetResult(&result).
		SetBody(map[string]interface{}{"stuid": stuid, "course_id": course_id}).
		ForceContentType("application/json").
		EnableTrace().
		Post(url)
	showTrace(resp)
	if result.Code != 200 {
		log.Printf("[HTTP] Response: %v\n", resp)
		return errors.New("响应码错误 - " + result.Msg)
	}
	return err
}

func DropCourse(url string, stuid string, course_id string) error {
	client := resty.New()
	var result CourseResult
	resp, err := client.R().
		SetResult(&result).
		SetBody(map[string]interface{}{"stuid": stuid, "course_id": course_id}).
		ForceContentType("application/json").
		EnableTrace().
		Post(url)
	showTrace(resp)
	if result.Code != 200 {
		log.Printf("[HTTP] Response: %v\n", resp)
		return errors.New("响应码错误 - " + result.Msg)
	}
	return err
}
