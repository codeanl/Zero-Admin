/* empty css             *//* empty css                  *//* empty css                    *//* empty css                *//* empty css                        *//* empty css                 *//* empty css                      *//* empty css               *//* empty css                  *//* empty css                  *//* empty css                        *//* empty css                      *//* empty css                   *//* empty css                  *//* empty css                *//* empty css                     */import{at as U,d as ue,a as d,r as ie,aa as pe,a7 as i,A as de,o as S,c as j,e,w as t,f as o,M as C,h as c,t as B,C as ce,F as me,K as D,k as fe,m as _e,aj as ge,ak as ve,n as Ve,p as ye,a5 as he,al as be,P as Ce,am as we,an as ke,ao as xe,ax as Ee,G as ze,ac as Ue,ap as Ae}from"./index-bc94ad6c.js";import{u as Se}from"./setting-343fd123.js";const Pe=(g,w,p,V)=>U.get("/api/sms/homeAdvertise/list",{params:{current:g,pageSize:w,name:p,status:V}}),G=g=>U.post("/api/sms/homeAdvertise/delete",g),M=g=>g.id?U.post("/api/sms/homeAdvertise/update",g):U.post("/api/sms/homeAdvertise/add",g),qe=["src"],Ie=["src"],Ne={style:{flex:"auto"}},$e=ue({__name:"index",setup(g){let w=Se(),p=d(1),V=d(10),P=d(0),k=d([]),y=d(""),h=d(""),x=d([]),A=d([]);const O=s=>{A.value=s};let r=d(),m=d(!1),n=ie({url:"",status:"",pic:"",note:"",name:"",clickCount:0,sort:0});pe(()=>{b()});const b=async(s=1)=>{p.value=s;let l=await Pe(p.value,V.value,y.value,h.value);l.code==200?(P.value=l.total,k.value=l.data,i({type:"success",message:l.message})):i({type:"error",message:l.message})},R=()=>{b()},T=()=>{b(),y.value="",h.value=""},F=()=>{w.refresh=!w.refresh},L=s=>{m.value=!0,Object.assign(n,s),D(()=>{r.value.clearValidate("name"),r.value.clearValidate("place"),r.value.clearValidate("status"),r.value.clearValidate("pic"),r.value.clearValidate("phone"),r.value.clearValidate("principal")})},H=()=>{m.value=!0,Object.assign(n,{id:0,url:"",status:"",pic:"",note:"",name:"",clickCount:0}),n.status="1",D(()=>{r.value.clearValidate("name"),r.value.clearValidate("place"),r.value.clearValidate("status"),r.value.clearValidate("pic"),r.value.clearValidate("phone"),r.value.clearValidate("principal")})},J=async s=>{x.value.push(s);const l={ids:x.value};let u=await G(l);u.code==200?(b(k.value.length>1?p.value:p.value-1),i({type:"success",message:u.message})):i({type:"error",message:u.message})},K=()=>{m.value=!1},Q=async()=>{let s=await M(n);s.code==200?(m.value=!1,i({type:"success",message:s.message}),b()):(m.value=!1,i({type:"error",message:s.message})),i({type:"success",message:s.message})},W=s=>{n.pic=s.data,r.value.clearValidate("pic")},X=s=>{if(s.type==="image/png"||s.type==="image/jpeg"||s.type==="image/gif"){if(s.size/1024/1024<4)return!0;i({type:"error",message:"上传的文件大小应小于4M"})}else return i({type:"error",message:"上传的文件类型必须是PNG|JPG|GIF"}),!1},Y=async()=>{x.value=A.value.map(u=>u.id);const s={ids:x.value};let l=await G(s);l.code===200?(b(k.value.length>1?p.value:p.value-1),i({type:"success",message:l.message})):i({type:"error",message:l.message})};let Z=async s=>{let l={id:s.id,status:s.status},u=await M(l);u.code===200?i({type:"success",message:u.message}):i({type:"error",message:u.message})};return(s,l)=>{const u=fe,f=_e,E=ge,q=ve,v=Ve,I=ye,N=he,_=be,$=Ce,ee=we,le=ke,ae=xe,te=Ee,se=de("Plus"),oe=ze,ne=Ue,re=Ae;return S(),j(me,null,[e(N,null,{default:t(()=>[e(I,{inline:!0},{default:t(()=>[e(f,{label:"名称:"},{default:t(()=>[e(u,{placeholder:"请输入搜索的用户名",modelValue:o(y),"onUpdate:modelValue":l[0]||(l[0]=a=>C(y)?y.value=a:y=a)},null,8,["modelValue"])]),_:1}),e(f,{label:"状态:"},{default:t(()=>[e(q,{modelValue:o(h),"onUpdate:modelValue":l[1]||(l[1]=a=>C(h)?h.value=a:h=a),class:"m-2",placeholder:"请选择状态"},{default:t(()=>[e(E,{label:"正常",value:"1"}),e(E,{label:"禁用",value:"0"})]),_:1},8,["modelValue"])]),_:1}),e(f,null,{default:t(()=>[e(v,{type:"primary",size:"default",onClick:T,disabled:!(o(y).length||o(h).length)},{default:t(()=>[c(" 搜索 ")]),_:1},8,["disabled"]),e(v,{size:"default",onClick:F},{default:t(()=>[c("重置")]),_:1})]),_:1})]),_:1})]),_:1}),e(N,null,{default:t(()=>[e(v,{type:"success",size:"default",onClick:H},{default:t(()=>[c(" 添加 ")]),_:1}),e(v,{type:"danger",size:"default",onClick:Y,disabled:!o(A).length},{default:t(()=>[c(" 批量删除 ")]),_:1},8,["disabled"]),e(le,{border:"",data:o(k),onSelectionChange:O,style:{margin:"15px 0"}},{default:t(()=>[e(_,{type:"selection",align:"center",width:"40px"}),e(_,{label:"id",align:"center",prop:"id",width:"50px"}),e(_,{label:"广告图片",align:"center",prop:"pic","show-overflow-tooltip":"",width:"120px"},{default:t(({row:a})=>[B("img",{src:a.pic,alt:"广告图片",style:{width:"100px",height:"auto"}},null,8,qe)]),_:1}),e(_,{label:"名称",align:"center",prop:"name","show-overflow-tooltip":""}),e(_,{label:"链接",align:"center",prop:"url","show-overflow-tooltip":""}),e(_,{label:"备注",align:"center",prop:"note","show-overflow-tooltip":""}),e(_,{label:"排序",align:"center",prop:"sort","show-overflow-tooltip":""}),e(_,{label:"状态",align:"center",prop:"status","show-overflow-tooltip":"",width:"60px"},{default:t(({row:a})=>[e($,{modelValue:a.status,"onUpdate:modelValue":z=>a.status=z,class:"ml-2",style:{"--el-switch-on-color":"#13ce66","--el-switch-off-color":"#ff4949"},"active-value":"1","inactive-value":"0",onChange:z=>o(Z)(a)},null,8,["modelValue","onUpdate:modelValue","onChange"])]),_:1}),e(_,{label:"操作",width:"300px",align:"center"},{default:t(({row:a})=>[e(v,{type:"primary",size:"small",icon:"Edit",onClick:z=>L(a)},{default:t(()=>[c(" 编辑 ")]),_:2},1032,["onClick"]),e(ee,{title:`你确定删除${a.name}`,width:"260px",onConfirm:z=>J(a.id)},{reference:t(()=>[e(v,{type:"danger",size:"small",icon:"Delete"},{default:t(()=>[c(" 删除 ")]),_:1})]),_:2},1032,["title","onConfirm"])]),_:1})]),_:1},8,["data"]),e(ae,{"current-page":o(p),"onUpdate:currentPage":l[2]||(l[2]=a=>C(p)?p.value=a:p=a),"page-size":o(V),"onUpdate:pageSize":l[3]||(l[3]=a=>C(V)?V.value=a:V=a),"page-sizes":[5,10,15,20],background:!0,layout:"prev, pager, next, jumper, -> , sizes, total",total:o(P),onCurrentChange:b,onSizeChange:R},null,8,["current-page","page-size","total"])]),_:1}),e(re,{modelValue:o(m),"onUpdate:modelValue":l[9]||(l[9]=a=>C(m)?m.value=a:m=a),title:o(n).id?"更新":"添加"},{footer:t(()=>[B("div",Ne,[e(v,{onClick:K},{default:t(()=>[c("取消")]),_:1}),e(v,{type:"primary",onClick:Q},{default:t(()=>[c("确定")]),_:1})])]),default:t(()=>[e(I,{model:o(n),ref_key:"formRef",ref:r},{default:t(()=>[e(f,{label:"名称",prop:"name"},{default:t(()=>[e(u,{placeholder:"请您输入名称",modelValue:o(n).name,"onUpdate:modelValue":l[4]||(l[4]=a=>o(n).name=a)},null,8,["modelValue"])]),_:1}),e(f,{label:"链接",prop:"url"},{default:t(()=>[e(u,{placeholder:"请您输入链接",modelValue:o(n).url,"onUpdate:modelValue":l[5]||(l[5]=a=>o(n).url=a)},null,8,["modelValue"])]),_:1}),e(f,{label:"备注",prop:"note"},{default:t(()=>[e(u,{placeholder:"请您输入备注",modelValue:o(n).note,"onUpdate:modelValue":l[6]||(l[6]=a=>o(n).note=a)},null,8,["modelValue"])]),_:1}),e(f,{label:"排序"},{default:t(()=>[c(" 排在第"),e(te,{modelValue:o(n).sort,"onUpdate:modelValue":l[7]||(l[7]=a=>o(n).sort=a),min:1,size:"small","controls-position":"right"},null,8,["modelValue"]),c("位 ")]),_:1}),e(f,{label:"状态",prop:"status"},{default:t(()=>[e(q,{modelValue:o(n).status,"onUpdate:modelValue":l[8]||(l[8]=a=>o(n).status=a),class:"m-2",placeholder:"请选择性别"},{default:t(()=>[e(E,{label:"不显示",value:"0"}),e(E,{label:"显示",value:"1"})]),_:1},8,["modelValue"])]),_:1}),e(f,{label:"广告图片",prop:"pic"},{default:t(()=>[e(ne,{class:"avatar-uploader",action:"/api/api/sys/upload","show-file-list":!1,"on-success":W,"before-upload":X},{default:t(()=>[o(n).pic?(S(),j("img",{key:0,src:o(n).pic,class:"avatar"},null,8,Ie)):(S(),ce(oe,{key:1,class:"avatar-uploader-icon"},{default:t(()=>[e(se)]),_:1}))]),_:1})]),_:1})]),_:1},8,["model"])]),_:1},8,["modelValue","title"])],64)}}});export{$e as default};