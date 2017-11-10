package json

// these 2 structs implement, by composition, the usecases.JsonW interface
// when there are more methods in the interface, you can create any structs to implement the interface
// with any method version you want
//
type JSONwriterStd struct {
	formatGetHistoriesRespStd
}

type JSONwriterSimple struct {
	formatGetHistoriesRespSimpleDates
}
