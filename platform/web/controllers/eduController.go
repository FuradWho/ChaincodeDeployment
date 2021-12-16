package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/FuradWho/ChaincodeDeployment/platform/models"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	log "github.com/sirupsen/logrus"
)

type EduController struct {
	Ctx     iris.Context
	Service services.ServiceSetup
}

var validate = validator.New()

func (c *EduController) GetTest() models.ResponseBean {
	path := c.Ctx.Path()
	log.Infoln(path)
	return models.SuccessMsg("1111")
}

func (c *EduController) SaveEdu() models.ResponseBean {

	path := c.Ctx.Path()
	log.Infoln(path)

	eduInfo := services.Education{
		ObjectType: c.Ctx.PostValueTrim("docType"),
		Name:       c.Ctx.PostValueTrim("Name"),        // 姓名
		Gender:     c.Ctx.PostValueTrim("Gender"),      // 性别
		Nation:     c.Ctx.PostValueTrim("userNation"),  // 民族
		EntityID:   c.Ctx.PostValueTrim("userIdCard"),  // 身份证号
		Place:      c.Ctx.PostValueTrim("userAddress"), // 籍贯
		BirthDay:   c.Ctx.PostValueTrim("BirthDay"),    // 出生日期

		EnrollDate:     c.Ctx.PostValueTrim("userEnterTime"),  // 入学日期
		GraduationDate: c.Ctx.PostValueTrim("GraduationDate"), // 毕（结）业日期
		SchoolName:     c.Ctx.PostValueTrim("userSchool"),     // 学校名称
		Major:          c.Ctx.PostValueTrim("userMajor"),      // 专业
		QuaType:        c.Ctx.PostValueTrim("QuaType"),        // 学历类别
		Length:         c.Ctx.PostValueTrim("Length"),         // 学制
		Mode:           c.Ctx.PostValueTrim("Mode"),           // 学习形式
		Level:          c.Ctx.PostValueTrim("Level"),          // 层次
		Graduation:     c.Ctx.PostValueTrim("Graduation"),     // 毕（结）业
		CertNo:         c.Ctx.PostValueTrim("CertNo"),         // 证书编号
	}

	//err := c.Ctx.ReadJSON(&eduInfo)
	//if err != nil {
	//	log.Errorf("Failed to json read to struct error : %s", err)
	//	return models.FailedMsg("Failed to json read to struct error")
	//}

	err := validate.Struct(eduInfo)
	if err != nil {
		log.Errorf("Failed to struct format error : %s", err)
		return models.FailedMsg("Failed to struct format error")

	}
	txId, err := c.Service.SaveEdu(eduInfo)
	if err != nil {
		log.Errorf("Failed to service save edu info : %s", err)
		return models.FailedMsg("Failed to service save edu info")
	}
	return models.SuccessData(map[string]string{
		"txId": txId,
	})

}

func (c *EduController) FindEduInfoByEntityID(entityID string) models.ResponseBean {

	path := c.Ctx.Path()
	log.Infoln(path)

	if entityID == "" {
		log.Errorf("Failed to Get Info because entity id is empty")
		return models.FailedMsg("Failed to Get Info because entity id is empty")
	}

	fmt.Println(entityID)
	result, err := c.Service.FindEduInfoByEntityID(entityID)
	if err != nil {
		log.Errorf("Failed to FindEduInfoByEntityID : %s", err)
		return models.FailedMsg("Failed to FindEduInfoByEntityID")
	}
	var eduInfo services.Education
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(result, &eduInfo)
	}
	return models.SuccessData(map[string]services.Education{
		"eduInfo": eduInfo,
	})
}

func (c *EduController) ModifyEdu() models.ResponseBean {
	path := c.Ctx.Path()
	log.Infoln(path)

	eduInfo := services.Education{
		ObjectType: c.Ctx.PostValueTrim("docType"),
		Name:       c.Ctx.PostValueTrim("Name"),        // 姓名
		Gender:     c.Ctx.PostValueTrim("Gender"),      // 性别
		Nation:     c.Ctx.PostValueTrim("userNation"),  // 民族
		EntityID:   c.Ctx.PostValueTrim("userIdCard"),  // 身份证号
		Place:      c.Ctx.PostValueTrim("userAddress"), // 籍贯
		BirthDay:   c.Ctx.PostValueTrim("BirthDay"),    // 出生日期

		EnrollDate:     c.Ctx.PostValueTrim("userEnterTime"),  // 入学日期
		GraduationDate: c.Ctx.PostValueTrim("GraduationDate"), // 毕（结）业日期
		SchoolName:     c.Ctx.PostValueTrim("userSchool"),     // 学校名称
		Major:          c.Ctx.PostValueTrim("userMajor"),      // 专业
		QuaType:        c.Ctx.PostValueTrim("QuaType"),        // 学历类别
		Length:         c.Ctx.PostValueTrim("Length"),         // 学制
		Mode:           c.Ctx.PostValueTrim("Mode"),           // 学习形式
		Level:          c.Ctx.PostValueTrim("Level"),          // 层次
		Graduation:     c.Ctx.PostValueTrim("Graduation"),     // 毕（结）业
		CertNo:         c.Ctx.PostValueTrim("CertNo"),         // 证书编号
	}

	//err := c.Ctx.ReadJSON(&eduInfo)
	//if err != nil {
	//	log.Errorf("Failed to json read to struct error : %s", err)
	//	return models.FailedMsg("Failed to json read to struct error")
	//}

	err := validate.Struct(eduInfo)
	if err != nil {
		log.Errorf("Failed to struct format error : %s", err)
		return models.FailedMsg("Failed to struct format error")

	}

	txId, err := c.Service.ModifyEdu(eduInfo)
	if err != nil {
		log.Errorf("Failed to struct ModifyEdu error : %s", err)
		return models.FailedMsg("Failed to ModifyEdu save edu info")
	}

	return models.SuccessData(map[string]string{
		"txId": txId,
	})
}

func (c *EduController) FindEduByCertNoAndName() models.ResponseBean {
	path := c.Ctx.Path()
	log.Infoln(path)

	certNo := c.Ctx.PostValueTrim("certNo")
	name := c.Ctx.PostValueTrim("name")

	if certNo == "" || name == "" {
		log.Errorf("Failed to param empty error")
		return models.FailedMsg("Failed to param empty error")
	}

	result, err := c.Service.FindEduByCertNoAndName(certNo, name)
	if err != nil {
		log.Errorf("Failed to Find Edu By CertNo And Name : %s", err)
		return models.FailedMsg("Failed to Find Edu By CertNo And Name")
	}

	var eduInfo services.Education
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(result, &eduInfo)
	}
	return models.SuccessData(map[string]services.Education{
		"eduInfo": eduInfo,
	})

}

func (c *EduController) BeforeActivation(b mvc.BeforeActivation) {

	b.Handle(
		"POST",
		"/save_edu",
		"SaveEdu",
	)
	b.Handle(
		"POST",
		"/modify_edu",
		"ModifyEdu",
	)

	b.Handle(
		"POST",
		"/find_by_certNoAndName",
		"FindEduByCertNoAndName",
	)

	b.Handle(
		"GET",
		"/find_by_entityID/{entityID:string}",
		"FindEduInfoByEntityID",
	)

}
