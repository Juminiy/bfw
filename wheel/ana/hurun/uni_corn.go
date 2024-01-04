package hurun

import (
	my_json "github.com/json-iterator/go"
	"os"
)

type (
	UniCorn struct {
		Rows  []UniCornElem `json:"rows"`
		Total int           `json:"total"`
	}
	UniCornElem struct {
		Wealth        float64 `json:"hs_Rank_Unicorn_Wealth"`
		FounderNameCN string  `json:"hs_Rank_Unicorn_ChaName_Cn"`
		FounderNameEN string  `json:"hs_Rank_Unicorn_ChaName_En"`
		ComNameCN     string  `json:"hs_Rank_Unicorn_ComName_Cn"`
		ComNameEN     string  `json:"hs_Rank_Unicorn_ComName_En"`
		ComLocCN      string  `json:"hs_Rank_Unicorn_ComHeadquarters_Cn"`
		ComLocEN      string  `json:"hs_Rank_Unicorn_ComHeadquarters_En"`
		ComAreaCN     string  `json:"hs_Rank_Unicorn_Industry_Cn"`
		ComAreaEN     string  `json:"hs_Rank_Unicorn_Industry_En"`
	}
)

func (u *UniCorn) TotalCNT() int {
	return u.Total
}

func (u *UniCorn) GetFounderNameList(top ...int) []string {
	var nameList []string
	destTop := 0xffffffff
	if len(top) > 0 && top[0] > 0 {
		destTop = top[0]
	}
	for i, elem := range u.Rows {
		if i+1 >= destTop {
			break
		}
		nameList = append(nameList, elem.FounderNameCN)
	}
	return nameList
}

func (u *UniCorn) GetCompanyNameList(top ...int) []string {
	var nameList []string
	destTop := 0xffffffff
	if len(top) > 0 && top[0] > 0 {
		destTop = top[0]
	}
	for i, elem := range u.Rows {
		if i+1 >= destTop {
			break
		}
		nameList = append(nameList, elem.ComNameCN)
	}
	return nameList
}

func GetUniCornEnterpriseTotalCount() int {
	h := MakeHurunURL()
	h.Num(EnterpriseUniCorn)
	h.Range(0, 0)
	var uniCornResp UniCorn
	h.Request(&uniCornResp)
	return uniCornResp.Total
}

func GetUniCornEnterpriseParse2JSONFile(outputFile string, search string, rank ...int) *UniCorn {
	h := MakeHurunURL()
	h.Num(EnterpriseUniCorn)
	if rankLen := len(rank); rankLen == 0 {

	} else if rankLen == 1 {
		h.Range(0, rank[0]-1)
	} else if rankLen == 2 {
		h.Range(rank[0], rank[1])
	} else {
		panic(requestHurunRankModeError)
	}
	if len(search) > 0 {
		h.Search(search)
	}
	var uniCornResp UniCorn
	h.Request(&uniCornResp)

	uniCornJsonStr, err := my_json.MarshalToString(uniCornResp)
	if err != nil {
		panic(err)
	}

	filePtr, err := os.OpenFile(getJSONFileName(hurunWriteDir, outputFile), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer func(filePtr *os.File) {
		err = filePtr.Close()
		if err != nil {
			panic(err)
		}
	}(filePtr)

	_, err = filePtr.WriteString(uniCornJsonStr)
	if err != nil {
		panic(err)
	}
	return &uniCornResp
}
