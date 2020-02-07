package types

//
// PermissionsClaim is a string to string slice
// mapping that should look as follows
// {
//	 "GET": []string{
//				"/pattern",
//				"/another/pattern",
//			},
//	 "POST": []string{
//				"/pattern",
//				"/another/pattern",
//			},
// }
//
type PermissionsClaim map[string][]string
