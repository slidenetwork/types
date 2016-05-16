package types

// type IStringEnum interface {
// 	Int() int
// 	String() string
// }

// // enum @DOC
// type StringEnum int

// // Value enum valuer
// func (e StringEnum) Value() (driver.Value, error) {
// 	return e.String(), nil
// }

// func (s *StringEnum) Scan(value interface{}) error {
// 	fromString(string(value.([]byte)))
// 	return nil
// }

// func (e StringEnum) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(e.String())
// }

// func (e *StringEnum) UnmarshalJSON(data []byte) error {
// 	var str string
// 	if err := json.Unmarshal(data, &str); err != nil {
// 		return err
// 	}
// 	e.fromString(str)
// 	return nil
// }

// func (e StringEnum) String(str string) StringEnum {
// 	*e = fromString(str, enumMap)
// }

// func fromString(str string, options map[int]string) int {
// 	for k, v := range options {
// 		if v == str {
// 			return k
// 		}
// 	}
// 	return 0
// }
