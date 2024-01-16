package reflect

type AStruct struct {
	AIntField       int         `bfw:"i"`
	AStringField    string      `bfw:"s"`
	AInterfaceField interface{} `bfw:"itf"`
}

func RefStruct() {

}
