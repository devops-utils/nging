//Do not edit this file, which is automatically generated by the generator.
package dbschema

import (
	"github.com/webx-top/db/lib/factory"
)

func init(){
	factory.Fields=map[string]map[string]*factory.FieldInfo{"ftp_user":map[string]*factory.FieldInfo{"ip_blacklist":&factory.FieldInfo{Name:"ip_blacklist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP黑名单(一行一个) ", GoType:"string", GoName:"IpBlacklist"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "group_id":&factory.FieldInfo{Name:"group_id", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"用户组", GoType:"uint", GoName:"GroupId"}, "password":&factory.FieldInfo{Name:"password", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:150, Options:[]string{}, DefaultValue:"", Comment:"密码", GoType:"string", GoName:"Password"}, "banned":&factory.FieldInfo{Name:"banned", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁止连接", GoType:"string", GoName:"Banned"}, "directory":&factory.FieldInfo{Name:"directory", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"授权目录(一行一个) ", GoType:"string", GoName:"Directory"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间 ", GoType:"uint", GoName:"Created"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "username":&factory.FieldInfo{Name:"username", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:120, Options:[]string{}, DefaultValue:"", Comment:"用户名", GoType:"string", GoName:"Username"}, "ip_whitelist":&factory.FieldInfo{Name:"ip_whitelist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP白名单(一行一个) ", GoType:"string", GoName:"IpWhitelist"}}, "ftp_user_group":map[string]*factory.FieldInfo{"id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"", GoType:"uint", GoName:"Id"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"修改时间", GoType:"uint", GoName:"Updated"}, "banned":&factory.FieldInfo{Name:"banned", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁止组内用户连接", GoType:"string", GoName:"Banned"}, "ip_whitelist":&factory.FieldInfo{Name:"ip_whitelist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP白名单(一行一个)", GoType:"string", GoName:"IpWhitelist"}, "ip_blacklist":&factory.FieldInfo{Name:"ip_blacklist", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"IP黑名单(一行一个)", GoType:"string", GoName:"IpBlacklist"}, "name":&factory.FieldInfo{Name:"name", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:255, Options:[]string{}, DefaultValue:"", Comment:"组名称", GoType:"string", GoName:"Name"}, "created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否禁用", GoType:"string", GoName:"Disabled"}, "directory":&factory.FieldInfo{Name:"directory", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"授权目录", GoType:"string", GoName:"Directory"}}, "vhost":map[string]*factory.FieldInfo{"created":&factory.FieldInfo{Name:"created", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"创建时间", GoType:"uint", GoName:"Created"}, "updated":&factory.FieldInfo{Name:"updated", DataType:"int", Unsigned:true, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"0", Comment:"更新时间", GoType:"uint", GoName:"Updated"}, "setting":&factory.FieldInfo{Name:"setting", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"设置", GoType:"string", GoName:"Setting"}, "disabled":&factory.FieldInfo{Name:"disabled", DataType:"enum", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{"Y", "N"}, DefaultValue:"N", Comment:"是否停用", GoType:"string", GoName:"Disabled"}, "id":&factory.FieldInfo{Name:"id", DataType:"int", Unsigned:true, PrimaryKey:true, AutoIncrement:true, Min:0, Max:0, Precision:0, MaxSize:10, Options:[]string{}, DefaultValue:"", Comment:"ID", GoType:"uint", GoName:"Id"}, "domain":&factory.FieldInfo{Name:"domain", DataType:"", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:0, Options:[]string{}, DefaultValue:"", Comment:"域名", GoType:"string", GoName:"Domain"}, "root":&factory.FieldInfo{Name:"root", DataType:"varchar", Unsigned:false, PrimaryKey:false, AutoIncrement:false, Min:0, Max:0, Precision:0, MaxSize:500, Options:[]string{}, DefaultValue:"", Comment:"网站物理路径", GoType:"string", GoName:"Root"}}}

}
