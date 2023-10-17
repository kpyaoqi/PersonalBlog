package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	// 设置Hexo博客的目录路径
	mdDir := "C:/Users/zhuba/Desktop/PersonalBlog/source/_posts"
	noteimgDir := "C:/Users/zhuba/Desktop/PersonalBlog/source/noteimg"
	// 设置Markdown文件所在的目录路径
	markdownDir := "C:/Users/zhuba/Desktop/笔记"
	err := generateHexoFiles(mdDir, noteimgDir, markdownDir)
	if err != nil {
		fmt.Println("生成Hexo文件失败:", err)
		return
	}
	fmt.Println("成功生成Hexo文件")
}

var date = time.Date(2021, 11, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02")

func generateHexoFiles(mdDir string, noteimgDir string, markdownDir string) error {
	files, err := ioutil.ReadDir(markdownDir)
	if err != nil {
		return fmt.Errorf("读取Markdown文件夹失败: %v", err)
	}
	for _, file := range files {
		filePath := filepath.Join(markdownDir, file.Name())
		if file.IsDir() {
			dirName := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			newmdDir := filepath.Join(mdDir, dirName)
			newnoteimgDir := filepath.Join(noteimgDir, dirName)
			err := os.MkdirAll(newmdDir, os.ModePerm)
			if err != nil {
				fmt.Printf("创建Hexo文件夹 %s 失败: %v\n", newmdDir, err)
				continue
			}
			err = generateHexoFiles(newmdDir, newnoteimgDir, filePath)
			if err != nil {
				fmt.Printf("生成Hexo文件失败: %v\n", err)
				continue
			}
		} else {
			if filepath.Ext(file.Name()) == ".md" {
				content, err := ioutil.ReadFile(filePath)
				if err != nil {
					fmt.Printf("读取Markdown文件 %s 失败: %v\n", file.Name(), err)
					continue
				}
				title := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
				categories := filepath.Base(markdownDir)
				tags := strings.ReplaceAll(strings.TrimPrefix(markdownDir, "C:\\Users\\zhuba\\Desktop\\笔记\\"), "\\", ",")
				re := regexp.MustCompile(`^[\d-]+`)
				tags = re.ReplaceAllString(tags, "")
				mdContent := fmt.Sprintf(`---
title: %s

date: %s	

categories: %s	

tags: [%s]
---	

%s`, title, date, categories, tags, content)
				addTime(&date)
				mdContent = strings.Replace(mdContent, markdownDir, mdDir, -1)
				mdFilePath := filepath.Join(mdDir, fmt.Sprintf("%s.md", title))
				_, err = os.Stat(filePath)
				if err == nil {
					break
				} else if os.IsNotExist(err) {
					//替换图片地址
					newString := strings.TrimPrefix(mdDir, "C:\\Users\\zhuba\\Desktop\\md\\")
					newString = strings.Replace(newString, "\\", "/", -1)
					mdContent = strings.ReplaceAll(mdContent, "(img/", "(/noteimg/"+newString+"/img/")
					mdContent = strings.ReplaceAll(mdContent, "src=\"img/", "src=\"/noteimg/"+newString+"/img/")
					err = ioutil.WriteFile(mdFilePath, []byte(mdContent), 0644)
					if err != nil {
						fmt.Printf("写入Hexo文件 %s 失败: %v\n", mdFilePath, err)
						continue
					}
				} else {
					fmt.Println("无法确定文件是否存在:", err)
				}
			} /*else {
				dstFilePath := filepath.Join(noteimgDir, file.Name())
				err := copyFile(filePath, dstFilePath)
				if err != nil {
					fmt.Printf("复制文件 %s 到 %s 失败: %v\n", filePath, dstFilePath, err)
					continue
				}
			}*/
		}
	}
	return nil
}

func copyFile(srcPath, dstPath string) error {
	// 打开源文件
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	// 确保目标目录存在，如果不存在则创建
	dstDir := filepath.Dir(dstPath)
	if _, err := os.Stat(dstDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dstDir, 0755); err != nil {
			return err
		}
	}
	// 创建目标文件
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	return nil
}

// 增加日期
func addTime(data *string) {
	// 将日期解析为时间对象
	t, _ := time.Parse("2006-01-02", date)
	b := rand.Intn(7) + rand.Intn(1)
	newDay := t.AddDate(0, 0, b)
	date = newDay.Format("2006-01-02")
}
