package main

//func main() {
//
//	//inc generic reset client
//	rc := &nrest.ReqType{
//		BaseUri: "0.0.0.0",
//		LogPath: "helpers.nrest",
//	}
//
//	//inc nube rest client
//	c := &nube_api.NubeRest{
//		Rest:          rc,
//		RubixPort:     nube_api.DefaultRubixService,
//		RubixUsername: "admin",
//		RubixPassword: "N00BWires",
//		UseRubixProxy: false,
//	}
//
//	//new nube rest client
//	nubeRest := nube_api.New(c)
//	nubeRest.GetToken()
//
//	//bacnet client
//	options := &nrest.ReqOpt{
//		Timeout:          500 * time.Second,
//		RetryCount:       1,
//		RetryWaitTime:    4 * time.Second,
//		RetryMaxWaitTime: 0,
//		Headers:          map[string]interface{}{"Authorization": nubeRest.RubixToken},
//	}
//	rc.LogPath = "helpers.nrest.bacnet.server"
//	rc.Port = nube_api.DefaultPortBacnet
//	c.RubixProxyPath = nube_api.ProxyBacnet
//	bacnetClient := &nube_api_bacnetserver.RestClient{
//		NubeRest: nubeRest,
//		Options:  options,
//	}
//	//get points
//	pnts, r := bacnetClient.GetPoints()
//	fmt.Println("ApiResponseIsBad", r.ApiReply.ApiResponseIsBad)
//	fmt.Println("ApiResponseIsJSON", r.ApiReply.ApiResponseIsJSON)
//	fmt.Println("ApiResponseLength", r.ApiReply.ApiResponseLength)
//	fmt.Println("Status code", r.ApiReply.Status())
//
//	if r.ApiReply.Err != nil {
//		fmt.Println("Error", r.Response)
//	} else {
//		for _, pnt := range pnts {
//			fmt.Println("points", pnt.ObjectType)
//		}
//		fmt.Println(pnts)
//	}
//
//	//get point
//	getPoint, r := bacnetClient.GetPoint("BhLtrFaNrtBxhVLyjc5CH")
//	fmt.Println("ApiResponseIsBad", r.ApiReply.ApiResponseIsBad)
//	fmt.Println("ApiResponseIsJSON", r.ApiReply.ApiResponseIsJSON)
//	fmt.Println("ApiResponseLength", r.ApiReply.ApiResponseLength)
//	fmt.Println("Status code", r.ApiReply.Status())
//	if r.ApiReply.Err != nil {
//		fmt.Println("Error", r.Response)
//	} else {
//		fmt.Println(getPoint)
//	}
//
//}
