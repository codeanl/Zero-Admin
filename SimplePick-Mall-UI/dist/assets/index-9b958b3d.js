/* empty css             *//* empty css               *//* empty css                     *//* empty css                 *//* empty css                    *//* empty css                *//* empty css                       *//* empty css                 */import{r as A}from"./index-0e4c8d23.js";import{d as N,r as P,a as B,A as F,f as b,c as p,e,w as l,t as C,p as G,j as w,o as r,C as m,h as g,a7 as V,a8 as S,a9 as M,m as j,k as z,G as q,ac as J,n as O,ai as T}from"./index-bc94ad6c.js";const H={key:0},K=C("h1",{style:{"text-align":"center","margin-top":"40px","font-size":"24px"}}," 商家｜自提点入驻 ",-1),L=["src"],Q=["src"],W=["src"],X={style:{float:"right"}},Y={key:1,class:"box"},de=N({__name:"index",setup(Z){const a=P({principalName:"",principalPhone:"",IDCardFront:"",IDCardReverse:"",name:"",address:"",type:"商户",pic:""});let f=B(!1);const x=t=>{a.pic=t.data},k=t=>{a.IDCardFront=t.data},I=t=>{a.IDCardReverse=t.data},i=t=>{if(t.type==="image/png"||t.type==="image/jpeg"||t.type==="image/gif"){if(t.size/1024/1024<4)return!0;V({type:"error",message:"上传的文件大小应小于4M"})}else return V({type:"error",message:"上传的文件类型必须是PNG|JPG|GIF"}),!1};let h=async()=>{a.type=="商户"&&(a.type="0"),a.type=="自提点"&&(a.type="1"),(await A(a)).code==200&&(f.value=!0)};return(t,o)=>{const y=S,E=M,n=j,d=z,c=F("Plus"),u=q,_=J,v=O,D=G,R=T,U=w;return b(f)==!1?(r(),p("div",H,[K,e(D,{model:a,"label-width":"120px",class:"box"},{default:l(()=>[e(n,{label:"类型"},{default:l(()=>[e(E,{modelValue:a.type,"onUpdate:modelValue":o[0]||(o[0]=s=>a.type=s)},{default:l(()=>[e(y,{label:"商户",value:"0"}),e(y,{label:"自提点",value:"1"})]),_:1},8,["modelValue"])]),_:1}),e(n,{label:"商户/自提点名称"},{default:l(()=>[e(d,{modelValue:a.name,"onUpdate:modelValue":o[1]||(o[1]=s=>a.name=s)},null,8,["modelValue"])]),_:1}),e(n,{label:"详细地址"},{default:l(()=>[e(d,{modelValue:a.address,"onUpdate:modelValue":o[2]||(o[2]=s=>a.address=s)},null,8,["modelValue"])]),_:1}),e(n,{label:"负责人姓名"},{default:l(()=>[e(d,{modelValue:a.principalName,"onUpdate:modelValue":o[3]||(o[3]=s=>a.principalName=s)},null,8,["modelValue"])]),_:1}),e(n,{label:"负责人联系方式"},{default:l(()=>[e(d,{modelValue:a.principalPhone,"onUpdate:modelValue":o[4]||(o[4]=s=>a.principalPhone=s)},null,8,["modelValue"])]),_:1}),e(n,{label:"门店图片",prop:"pic"},{default:l(()=>[e(_,{class:"avatar-uploader",action:"/api/api/sys/upload","show-file-list":!1,"on-success":x,"before-upload":i},{default:l(()=>[a.pic?(r(),p("img",{key:0,src:a.pic,class:"avatar"},null,8,L)):(r(),m(u,{key:1,class:"avatar-uploader-icon"},{default:l(()=>[e(c)]),_:1}))]),_:1})]),_:1}),e(n,{label:"身份证正面照",prop:"IDCardFront"},{default:l(()=>[e(_,{class:"avatar-uploader",action:"/api/api/sys/upload","show-file-list":!1,"on-success":k,"before-upload":i},{default:l(()=>[a.IDCardFront?(r(),p("img",{key:0,src:a.IDCardFront,class:"avatar"},null,8,Q)):(r(),m(u,{key:1,class:"avatar-uploader-icon"},{default:l(()=>[e(c)]),_:1}))]),_:1})]),_:1}),e(n,{label:"身份证反面照",prop:"IDCardReverse"},{default:l(()=>[e(_,{class:"avatar-uploader",action:"/api/api/sys/upload","show-file-list":!1,"on-success":I,"before-upload":i},{default:l(()=>[a.IDCardReverse?(r(),p("img",{key:0,src:a.IDCardReverse,class:"avatar"},null,8,W)):(r(),m(u,{key:1,class:"avatar-uploader-icon"},{default:l(()=>[e(c)]),_:1}))]),_:1})]),_:1}),C("div",X,[e(v,{type:"primary",onClick:b(h)},{default:l(()=>[g("提交申请")]),_:1},8,["onClick"])])]),_:1},8,["model"])])):(r(),p("div",Y,[e(U,null,{default:l(()=>[e(R,{icon:"success",title:"申请成功","sub-title":"您已申请成功，等待管理员审核！审核结果将通过手机短信发送至您的手机！"},{extra:l(()=>[e(v,{type:"primary"},{default:l(()=>[g("查看进度")]),_:1})]),_:1})]),_:1})]))}}});export{de as default};