/* empty css             *//* empty css                  *//* empty css                 *//* empty css                  *//* empty css                      *//* empty css               *//* empty css                  *//* empty css                        *//* empty css                      *//* empty css                   *//* empty css                *//* empty css                     */import{at as Y,d as te,a as f,r as oe,aa as ne,a7 as b,o as k,c as se,e,w as t,f as n,M as x,h as d,C,z as U,t as q,F as ue,K as H,k as pe,m as re,aj as de,ak as ie,n as me,p as ce,a5 as fe,al as _e,ah as ge,am as ve,an as be,ao as ye,az as Ve,ap as Te,_ as he}from"./index-bc94ad6c.js";import{u as we}from"./setting-343fd123.js";const ke=(g,z,m,y,P)=>Y.get("/api/sms/coupon/list",{params:{current:g,pageSize:z,type:m,name:y,useType:P}}),N=g=>Y.post("/api/sms/coupon/delete",g),Ee=g=>g.id?Y.post("/api/sms/coupon/update",g):Y.post("/api/sms/coupon/add",g),xe={class:"block"},Ce={class:"block"},Ue={style:{flex:"auto"}},ze=te({__name:"index",setup(g){let z=we(),m=f(1),y=f(10),P=f(0),D=f([]),V=f(""),T=f(""),h=f(""),L=f([]),M=f([]);const j=p=>{M.value=p};let i=f(),_=f(!1),o=oe({type:"",name:"",count:"",amount:"",perLimit:"",minPoint:"",startTime:"",endTime:"",enableTime:"",useType:"",note:"",code:"",starEndTime:[]});ne(()=>{w()});const w=async(p=1)=>{m.value=p;let a=await ke(m.value,y.value,T.value,V.value,h.value);a.code==200?(P.value=a.total,D.value=a.data,b({type:"success",message:a.message})):b({type:"error",message:a.message})},O=()=>{w()},R=()=>{_.value=!1},A=async()=>{o.amount=parseFloat(o.amount),o.minPoint=parseFloat(o.minPoint),o.count=parseInt(o.count),o.perLimit=parseInt(o.perLimit),o.startTime=o.starEndTime[0],o.endTime=o.starEndTime[1];let p=await Ee(o);p.code==200?(_.value=!1,b({type:"success",message:p.message}),w()):(_.value=!1,b({type:"error",message:p.message}))},K=()=>{_.value=!0,Object.assign(o,{id:0,type:"",name:"",count:"",amount:"",perLimit:"",minPoint:"",startTime:"",endTime:"",enableTime:"",useType:"",note:"",code:"",starEndTime:[]}),H(()=>{i.value.clearValidate("name"),i.value.clearValidate("place"),i.value.clearValidate("status"),i.value.clearValidate("pic"),i.value.clearValidate("phone"),i.value.clearValidate("principal")})},G=p=>{_.value=!0,Object.assign(o,p),o.starEndTime.push(o.startTime),o.starEndTime.push(o.endTime),H(()=>{i.value.clearValidate("name"),i.value.clearValidate("place"),i.value.clearValidate("status"),i.value.clearValidate("pic"),i.value.clearValidate("phone"),i.value.clearValidate("principal")})},J=async p=>{L.value.push(p);const a={ids:L.value};let r=await N(a);r.code==200?(w(D.value.length>1?m.value:m.value-1),b({type:"success",message:r.message})):b({type:"error",message:r.message})},Q=async()=>{L.value=M.value.map(r=>r.id);const p={ids:L.value};let a=await N(p);a.code===200?(w(D.value.length>1?m.value:m.value-1),b({type:"success",message:a.message})):b({type:"error",message:a.message})},W=()=>{z.refresh=!z.refresh},X=()=>{w(),V.value="",T.value="",h.value=""};return(p,a)=>{const r=pe,s=re,c=de,S=ie,v=me,F=ce,I=fe,u=_e,E=ge,Z=ve,$=be,ee=ye,B=Ve,le=Te;return k(),se(ue,null,[e(I,null,{default:t(()=>[e(F,{inline:!0},{default:t(()=>[e(s,{label:"名称:"},{default:t(()=>[e(r,{placeholder:"请输入搜索的用户名",modelValue:n(V),"onUpdate:modelValue":a[0]||(a[0]=l=>x(V)?V.value=l:V=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"优惠券类型:"},{default:t(()=>[e(S,{modelValue:n(T),"onUpdate:modelValue":a[1]||(a[1]=l=>x(T)?T.value=l:T=l),class:"m-2",placeholder:"请选择状态"},{default:t(()=>[e(c,{label:"全场赠券",value:"0"}),e(c,{label:"购物赠券",value:"1"}),e(c,{label:"注册赠券",value:"2"})]),_:1},8,["modelValue"])]),_:1}),e(s,{label:"使用类型:"},{default:t(()=>[e(S,{modelValue:n(h),"onUpdate:modelValue":a[2]||(a[2]=l=>x(h)?h.value=l:h=l),class:"m-2",placeholder:"请选择状态"},{default:t(()=>[e(c,{label:"全场通用",value:"0"}),e(c,{label:"指定分类",value:"1"}),e(c,{label:"指定商品",value:"2"})]),_:1},8,["modelValue"])]),_:1}),e(s,null,{default:t(()=>[e(v,{type:"primary",size:"default",onClick:X,disabled:!(n(V).length||n(T).length||n(h).length)},{default:t(()=>[d(" 搜索 ")]),_:1},8,["disabled"]),e(v,{size:"default",onClick:W},{default:t(()=>[d("重置")]),_:1})]),_:1})]),_:1})]),_:1}),e(I,null,{default:t(()=>[e(v,{type:"success",size:"default",onClick:K},{default:t(()=>[d(" 添加 ")]),_:1}),e(v,{type:"danger",size:"default",onClick:Q,disabled:!n(M).length},{default:t(()=>[d(" 批量删除 ")]),_:1},8,["disabled"]),e($,{border:"",data:n(D),onSelectionChange:j,style:{margin:"15px 0"}},{default:t(()=>[e(u,{type:"selection",align:"center",width:"40px"}),e(u,{label:"id",align:"center",prop:"id",width:"50px"}),e(u,{label:"类型",align:"center",prop:"type","show-overflow-tooltip":""},{default:t(({row:l})=>[l.type==="0"?(k(),C(E,{key:"item.label",class:"mx-1",type:"danger",effect:"light"},{default:t(()=>[d(" 全场赠券 ")]),_:1})):U("",!0),l.type==="1"?(k(),C(E,{key:"item.label",class:"mx-1",type:"success",effect:"light"},{default:t(()=>[d(" 购物赠券 ")]),_:1})):U("",!0),l.type==="2"?(k(),C(E,{key:"item.label",class:"mx-1",type:"danger",effect:"light"},{default:t(()=>[d(" 注册赠券 ")]),_:1})):U("",!0)]),_:1}),e(u,{label:"使用类型",align:"center",prop:"useType","show-overflow-tooltip":""},{default:t(({row:l})=>[l.useType==="0"?(k(),C(E,{key:"item.label",class:"mx-1",type:"danger",effect:"light"},{default:t(()=>[d(" 全场通用 ")]),_:1})):U("",!0),l.useType==="1"?(k(),C(E,{key:"item.label",class:"mx-1",type:"success",effect:"light"},{default:t(()=>[d(" 指定分类 ")]),_:1})):U("",!0),l.useType==="2"?(k(),C(E,{key:"item.label",class:"mx-1",type:"danger",effect:"light"},{default:t(()=>[d(" 指定商品 ")]),_:1})):U("",!0)]),_:1}),e(u,{label:"名称",align:"center",prop:"name","show-overflow-tooltip":""}),e(u,{label:"数量",align:"center",prop:"count","show-overflow-tooltip":""}),e(u,{label:"金额",align:"center",prop:"amount","show-overflow-tooltip":""}),e(u,{label:"每人限领张数",align:"center",prop:"perLimit","show-overflow-tooltip":""}),e(u,{label:"使用门槛",align:"center",prop:"minPoint","show-overflow-tooltip":""}),e(u,{label:"开始时间",align:"center",prop:"startTime","show-overflow-tooltip":""}),e(u,{label:"结束时间",align:"center",prop:"endTime","show-overflow-tooltip":""}),e(u,{label:"开始领取时间",align:"center",prop:"enableTime","show-overflow-tooltip":""}),e(u,{label:"备注",align:"center",prop:"note","show-overflow-tooltip":""}),e(u,{label:"优惠码",align:"center",prop:"code","show-overflow-tooltip":""}),e(u,{label:"操作",width:"200px",align:"center"},{default:t(({row:l})=>[e(v,{type:"primary",size:"small",icon:"Edit",onClick:ae=>G(l)},{default:t(()=>[d(" 编辑 ")]),_:2},1032,["onClick"]),e(Z,{title:`你确定删除${l.name}`,width:"260px",onConfirm:ae=>J(l.id)},{reference:t(()=>[e(v,{type:"danger",size:"small",icon:"Delete"},{default:t(()=>[d(" 删除 ")]),_:1})]),_:2},1032,["title","onConfirm"])]),_:1})]),_:1},8,["data"]),e(ee,{"current-page":n(m),"onUpdate:currentPage":a[3]||(a[3]=l=>x(m)?m.value=l:m=l),"page-size":n(y),"onUpdate:pageSize":a[4]||(a[4]=l=>x(y)?y.value=l:y=l),"page-sizes":[5,10,15,20],background:!0,layout:"prev, pager, next, jumper, -> , sizes, total",total:n(P),onCurrentChange:w,onSizeChange:O},null,8,["current-page","page-size","total"])]),_:1}),e(le,{modelValue:n(_),"onUpdate:modelValue":a[16]||(a[16]=l=>x(_)?_.value=l:_=l),title:n(o).id?"更新":"添加"},{footer:t(()=>[q("div",Ue,[e(v,{onClick:R},{default:t(()=>[d("取消")]),_:1}),e(v,{type:"primary",onClick:A},{default:t(()=>[d("确定")]),_:1})])]),default:t(()=>[e(F,{model:n(o),ref_key:"formRef",ref:i},{default:t(()=>[e(s,{label:"名称",prop:"name"},{default:t(()=>[e(r,{placeholder:"请您输入名称",modelValue:n(o).name,"onUpdate:modelValue":a[5]||(a[5]=l=>n(o).name=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"类型",prop:"type"},{default:t(()=>[e(S,{modelValue:n(o).type,"onUpdate:modelValue":a[6]||(a[6]=l=>n(o).type=l),class:"m-2",placeholder:"请选择性别"},{default:t(()=>[e(c,{label:"全场赠券",value:"0"}),e(c,{label:"购物赠券",value:"1"}),e(c,{label:"注册赠券",value:"2"})]),_:1},8,["modelValue"])]),_:1}),e(s,{label:"使用类型",prop:"useType"},{default:t(()=>[e(S,{modelValue:n(o).useType,"onUpdate:modelValue":a[7]||(a[7]=l=>n(o).useType=l),class:"m-2",placeholder:"请选择使用类型"},{default:t(()=>[e(c,{label:"全场通用",value:"0"}),e(c,{label:"指定分类",value:"1"}),e(c,{label:"指定商品",value:"2"})]),_:1},8,["modelValue"])]),_:1}),e(s,{label:"数量",prop:"count"},{default:t(()=>[e(r,{placeholder:"请您输入数量",modelValue:n(o).count,"onUpdate:modelValue":a[8]||(a[8]=l=>n(o).count=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"金额",prop:"amount"},{default:t(()=>[e(r,{placeholder:"请您输入金额",modelValue:n(o).amount,"onUpdate:modelValue":a[9]||(a[9]=l=>n(o).amount=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"每人限领张数",prop:"perLimit"},{default:t(()=>[e(r,{placeholder:"请您输入每人限领张数",modelValue:n(o).perLimit,"onUpdate:modelValue":a[10]||(a[10]=l=>n(o).perLimit=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"使用门槛",prop:"minPoint"},{default:t(()=>[e(r,{placeholder:"请您输入使用门槛",modelValue:n(o).minPoint,"onUpdate:modelValue":a[11]||(a[11]=l=>n(o).minPoint=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"起止时间",prop:"endTime"},{default:t(()=>[q("div",xe,[e(B,{modelValue:n(o).starEndTime,"onUpdate:modelValue":a[12]||(a[12]=l=>n(o).starEndTime=l),type:"datetimerange","range-separator":"到","start-placeholder":"开始时间","end-placeholder":"结束时间","value-format":"YYYY-MM-DD HH:mm:ss"},null,8,["modelValue"])])]),_:1}),e(s,{label:"领取时间",prop:"enableTime"},{default:t(()=>[q("div",Ce,[e(B,{modelValue:n(o).enableTime,"onUpdate:modelValue":a[13]||(a[13]=l=>n(o).enableTime=l),type:"datetime",placeholder:"请您输入开始领取时间","value-format":"YYYY-MM-DD HH:mm:ss"},null,8,["modelValue"])])]),_:1}),e(s,{label:"备注",prop:"note"},{default:t(()=>[e(r,{placeholder:"请您输入备注",modelValue:n(o).note,"onUpdate:modelValue":a[14]||(a[14]=l=>n(o).note=l)},null,8,["modelValue"])]),_:1}),e(s,{label:"优惠码",prop:"code"},{default:t(()=>[e(r,{placeholder:"请您输入优惠码",modelValue:n(o).code,"onUpdate:modelValue":a[15]||(a[15]=l=>n(o).code=l)},null,8,["modelValue"])]),_:1})]),_:1},8,["model"])]),_:1},8,["modelValue","title"])],64)}}});const Re=he(ze,[["__scopeId","data-v-13d78829"]]);export{Re as default};