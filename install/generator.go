package install

import (
	"bytes"
	"github.com/wonderivan/logger"
	"strings"
	"text/template"
)

func DockerCompose(nodeName string) string {
	var sb strings.Builder
	sb.Write([]byte(dockerCompose))
	var envMap = make(map[string]interface{})
	envMap["SERVER_URL"] = ServerUrl
	envMap["NODE_NAME"] = nodeName
	envMap["MANAGE_IMAGE"] = ManageImage
	return FromTemplateContent(sb.String(), envMap)
}

func BaseUtils(ip string) string {
	var sb strings.Builder
	sb.Write([]byte(baseUtils))
	var envMap = make(map[string]interface{})
	envMap["HostIp"] = ip
	return FromTemplateContent(sb.String(), envMap)
}

func FromTemplateContent(templateContent string, envMap map[string]interface{}) string {
	tmpl, err := template.New("text").Parse(templateContent)
	defer func() {
		if r := recover(); r != nil {
			logger.Error("Template parse failed:", err)
		}
	}()
	if err != nil {
		panic(1)
	}
	var buffer bytes.Buffer
	_ = tmpl.Execute(&buffer, envMap)
	return string(buffer.Bytes())
}