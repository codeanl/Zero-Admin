/* empty css             *//* empty css                  *//* empty css                        *//* empty css                 *//* empty css                      *//* empty css               *//* empty css                  *//* empty css                  *//* empty css                        *//* empty css                      *//* empty css                   *//* empty css                *//* empty css                     *//* empty css                */import{r as X,a as Y,b as Z}from"./index-34640b16.js";import{r as $}from"./index-b9f34230.js";import{u as ee}from"./setting-343fd123.js";import{d as ae,a as u,r as le,aa as te,a7 as b,o as E,c as oe,e,w as l,f as o,M as v,h as r,C as q,z as B,t as se,F as ne,k as re,m as ue,aw as de,aj as pe,ak as ie,n as me,p as ce,a5 as _e,al as fe,ah as ge,am as ye,an as be,ao as ve,ax as Ve,ap as Ce,_ as ke}from"./index-bc94ad6c.js";const xe={class:"dialog-footer"},Ee=ae({__name:"index",setup(ze){let z=u([]),V=u(1),C=u(10),I=u(0),i=u(""),m=u(""),f=u(),p=u(!1),x=u([]),U=u([]),s=le({id:0,name:"",sort:0,type:"",value:"",attributeCategoryID:0});te(()=>{g()});const g=async()=>{let n=await X(V.value,C.value,i.value,m.value,f.value);n.code===200&&(z.value=n.data,I.value=n.total);let t=await $();t.code===200&&(x.value=t.data),b({type:"success",message:n.message}),(n.code===400||t.code===400)&&b({type:"success",message:n.message})},P=()=>{g()};let T=()=>{Object.assign(s,{id:0,name:"",sort:0,type:"",value:"",attributeCategoryID:""}),p.value=!0},h=n=>{p.value=!0,Object.assign(s,JSON.parse(JSON.stringify(n)))},j=()=>{p.value=!1};const F=async()=>{s.sort=parseInt(s.sort);let n=await Y(s);n.code===200?(p.value=!1,g(),b({type:"success",message:n.message})):b({type:"error",message:n.message})};let L=async n=>{U.value.push(n);const t={ids:U.value};let c=await Z(t);c.code===200?(g(),b({type:"success",message:c.message})):b({type:"error",message:c.message})};const M=()=>{g(),i.value="",m.value=""};let S=ee();const J=()=>{S.refresh=!S.refresh};return(n,t)=>{const c=re,d=ue,D=de,k=pe,w=ie,_=me,N=ce,A=_e,y=fe,O=ge,R=ye,H=be,G=ve,K=Ve,Q=Ce;return E(),oe(ne,null,[e(A,null,{default:l(()=>[e(N,{inline:!0},{default:l(()=>[e(d,{label:"名称:"},{default:l(()=>[e(c,{placeholder:"请输入搜索的用户名",modelValue:o(i),"onUpdate:modelValue":t[0]||(t[0]=a=>v(i)?i.value=a:i=a)},null,8,["modelValue"])]),_:1}),e(d,{label:"所属属性:"},{default:l(()=>[e(D,{modelValue:o(f),"onUpdate:modelValue":t[1]||(t[1]=a=>v(f)?f.value=a:f=a),data:o(x),"check-strictly":"",props:{key:"categoryId",label:"name"},"node-key":"id","render-after-expand":!1},null,8,["modelValue","data"])]),_:1}),e(d,{label:"类型:"},{default:l(()=>[e(w,{modelValue:o(m),"onUpdate:modelValue":t[2]||(t[2]=a=>v(m)?m.value=a:m=a),class:"m-2",placeholder:"请选择类型"},{default:l(()=>[e(k,{label:"属性",value:"1"}),e(k,{label:"规格",value:"2"})]),_:1},8,["modelValue"])]),_:1}),e(d,null,{default:l(()=>[e(_,{type:"primary",size:"default",onClick:M,disabled:!(o(i).length||o(m).length||o(f))},{default:l(()=>[r(" 搜索 ")]),_:1},8,["disabled"]),e(_,{size:"default",onClick:J},{default:l(()=>[r("重置")]),_:1})]),_:1})]),_:1})]),_:1}),e(A,null,{default:l(()=>[e(_,{type:"success",size:"default",icon:"Plus",onClick:o(T)},{default:l(()=>[r("添加")]),_:1},8,["onClick"]),e(H,{border:"",style:{margin:"15px 0"},data:o(z)},{default:l(()=>[e(y,{label:"编号",width:"50px",align:"center",prop:"id"}),e(y,{label:"属性名称",align:"center",prop:"name"}),e(y,{label:"类型",align:"center",prop:"type"},{default:l(({row:a})=>[a.type==="1"?(E(),q(O,{key:0,class:"mx-1",type:"success",effect:"light"},{default:l(()=>[r(" 属性 ")]),_:1})):B("",!0),a.type==="2"?(E(),q(O,{key:1,class:"mx-1",type:"warning",effect:"light"},{default:l(()=>[r(" 规格 ")]),_:1})):B("",!0)]),_:1}),e(y,{label:"排序",align:"center",prop:"sort"}),e(y,{label:"属性值",align:"center",prop:"value"}),e(y,{label:"操作",align:"center"},{default:l(({row:a})=>[e(_,{type:"primary",size:"small",onClick:W=>o(h)(a),icon:"Edit"},{default:l(()=>[r("编辑")]),_:2},1032,["onClick"]),e(R,{title:`你确定删除${a.name}?`,width:"200px",onConfirm:W=>o(L)(a.id)},{reference:l(()=>[e(_,{type:"danger",size:"small",icon:"Delete"},{default:l(()=>[r("删除")]),_:1})]),_:2},1032,["title","onConfirm"])]),_:1})]),_:1},8,["data"]),e(G,{"current-page":o(V),"onUpdate:currentPage":t[3]||(t[3]=a=>v(V)?V.value=a:V=a),"page-size":o(C),"onUpdate:pageSize":t[4]||(t[4]=a=>v(C)?C.value=a:C=a),"page-sizes":[5,10,15,20],background:!0,layout:"prev, pager, next, jumper, -> , sizes, total",total:o(I),onCurrentChange:g,onSizeChange:P},null,8,["current-page","page-size","total"])]),_:1}),e(Q,{modelValue:o(p),"onUpdate:modelValue":t[10]||(t[10]=a=>v(p)?p.value=a:p=a)},{footer:l(()=>[se("span",xe,[e(_,{type:"primary",size:"default",onClick:F},{default:l(()=>[r("保存")]),_:1}),e(_,{type:"primary",size:"default",onClick:o(j)},{default:l(()=>[r("取消")]),_:1},8,["onClick"])])]),default:l(()=>[e(N,null,{default:l(()=>[e(d,{label:"所属分类"},{default:l(()=>[e(D,{modelValue:o(s).attributeCategoryID,"onUpdate:modelValue":t[5]||(t[5]=a=>o(s).attributeCategoryID=a),data:o(x),"check-strictly":"",props:{key:"attributeCategoryID",label:"name"},"node-key":"id","render-after-expand":!1},null,8,["modelValue","data"])]),_:1}),e(d,{label:"属性名称"},{default:l(()=>[e(c,{placeholder:"请输入属性名称",modelValue:o(s).name,"onUpdate:modelValue":t[6]||(t[6]=a=>o(s).name=a)},null,8,["modelValue"])]),_:1}),e(d,{label:"属性值"},{default:l(()=>[e(c,{placeholder:"请输入属性值",modelValue:o(s).value,"onUpdate:modelValue":t[7]||(t[7]=a=>o(s).value=a)},null,8,["modelValue"])]),_:1}),e(d,{label:"排序"},{default:l(()=>[r(" 排在第"),e(K,{modelValue:o(s).sort,"onUpdate:modelValue":t[8]||(t[8]=a=>o(s).sort=a),min:1,size:"small","controls-position":"right"},null,8,["modelValue"]),r("位 ")]),_:1}),e(d,{label:"类型"},{default:l(()=>[e(w,{modelValue:o(s).type,"onUpdate:modelValue":t[9]||(t[9]=a=>o(s).type=a),class:"m-2",placeholder:"请选择类型"},{default:l(()=>[e(k,{label:"属性",value:"1"}),e(k,{label:"规格",value:"2"})]),_:1},8,["modelValue"])]),_:1})]),_:1})]),_:1},8,["modelValue"])],64)}}});const Re=ke(Ee,[["__scopeId","data-v-0232a882"]]);export{Re as default};