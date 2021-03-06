package test

import (
	"fmt"
	"gengine/base"
	"gengine/builder"
	"gengine/context"
	"gengine/engine"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

type Person struct {
	Name string
	Age int64
}

func getPerson(n string, a int64) *Person{
	return &Person{
		Name: n,
		Age:  a,
	}
}

func Test_Struct(t *testing.T)  {
	 p := getPerson("777", 5)
	 println(p.Age)
}


const rule_s = `
rule "test_struct_return" "test" 
begin
	p = getPerson("hello", 100)
	Sout(p.Age)
end
`

func exe_struct(){
	/**
	不要注入除函数和结构体指针以外的其他类型(如变量)
	*/
	dataContext := context.NewDataContext()
	//注入结构体指针
	//重命名函数,并注入
	dataContext.Add("Sout",fmt.Println)
	dataContext.Add("getPerson",getPerson)
	//初始化规则引擎
	knowledgeContext := base.NewKnowledgeContext()
	ruleBuilder := builder.NewRuleBuilder(knowledgeContext, dataContext)

	//读取规则
	start1 := time.Now().UnixNano()
	err := ruleBuilder.BuildRuleFromString(rule_s)
	end1 := time.Now().UnixNano()

	logrus.Infof("规则个数:%d, 加载规则耗时:%d ns", len(knowledgeContext.RuleEntities), end1-start1 )

	if err != nil{
		logrus.Errorf("err:%s ", err)
	}else{
		eng := engine.NewGengine()

		start := time.Now().UnixNano()
		// true: means when there are many rules， if one rule execute error，continue to execute rules after the occur error rule
		err := eng.Execute(ruleBuilder, true)
		end := time.Now().UnixNano()
		if err != nil{
			logrus.Errorf("execute rule error: %v", err)
		}
		logrus.Infof("execute rule cost %d ns",end-start)
	}
}

func Test_struct(t *testing.T)  {
	exe_struct()
}






