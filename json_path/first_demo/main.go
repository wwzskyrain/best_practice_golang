package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

const iosPvJson = `[
    {
        "session_id": "BFFF2239-8DB5-4115-A23C-2A125D01F133",
        "unique_key": "1763821880931866_1735117458825_10000001_session_log",
        "pv_filters": {
            "sdk_tag": "PAGSDK",
            "app_version": "6.5.0.2",
            "bundle_id": "com.bytedance.ADUnion.Development",
            "sdk_version_name": "6.5.0.2",
            "sdk_state": "testFlight",
            "country_name": "CN",
            "os_version": "14.2",
            "backtrace_switch": "1",
            "device_model": "iPhone10,3",
            "app_name": "网盟测试",
            "device_type": "iPhone"
        },
        "device_id": "1763821880931866",
        "app_id": "10000001",
        "timestamp": 1735117458825
    },
    {
        "session_id": "10B9DE64-5F28-46C3-B3E3-F62B8D7C34B6",
        "unique_key": "1763821880931866_1735117524486_10000001_session_log",
        "pv_filters": {
            "sdk_tag": "PAGSDK",
            "app_version": "6.5.0.2",
            "bundle_id": "com.bytedance.ADUnion.Development",
            "sdk_version_name": "6.5.0.2",
            "sdk_state": "testFlight",
            "country_name": "CN",
            "os_version": "14.2",
            "backtrace_switch": "1",
            "device_model": "iPhone10,3",
            "app_name": "网盟测试",
            "device_type": "iPhone"
        },
        "device_id": "1763821880931866",
        "app_id": "10000001",
        "timestamp": 1735117524486
    }
]`

func main() {
	value := gjson.Get(json, "name.last")
	println(value.String())
	result := gjson.Parse(json)
	println(result.Type)
	arrayResult := result.Array()
	println(len(arrayResult))
	println(arrayResult[0].Get("name.last").String())

	r := gjson.Parse(iosPvJson)
	rr := r.Array()
	println(len(rr))

}
