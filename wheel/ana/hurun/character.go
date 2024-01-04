package hurun

type (
	Character struct {
		Gender        string `json:"hs_Character_Gender"`
		Birth         string `json:"hs_Character_Birthday"`
		NameCN        string `json:"hs_Character_Fullname_Cn"`
		NameEN        string `json:"hs_Character_Fullname_En"`
		FaceLink      string `json:"hs_Character_Photo"`
		NativePlaceCN string `json:"hs_Character_NativePlace_Cn"`
		NativePlaceEN string `json:"hs_Character_NativePlace_En"`
		Education     string `json:"hs_Character_Education_Cn"`
		School        string `json:"hs_Character_School_Cn"`
		Major         string `json:"hs_Character_Major_En"`
	}
)
