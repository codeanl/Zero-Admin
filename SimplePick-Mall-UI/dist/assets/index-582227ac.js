/* empty css             *//* empty css               *//* empty css                     *//* empty css                 *//* empty css               */import{d as k,r as v,a as b,u as I,b as F,c as S,e,w as a,E as B,o as C,f as s,g as N,l as q,h as H,i,j as U,k as $,m as R,n as T,p as j,q as z,s as L,t as f,_ as A}from"./index-bc94ad6c.js";import{g as p}from"./time-749b2ba7.js";const g=r=>(z("data-v-4a76de16"),r=r(),L(),r),D={class:"login_container"},G=g(()=>f("h1",null,"欢迎您",-1)),J=g(()=>f("h2",null,"感谢使用定点帮扶农产品直销系统",-1)),K=k({__name:"index",setup(r){let t=v({username:"admin",password:"123456"}),u=b();const w={username:[{required:!0,message:"请输入登陆名称",trigger:"blur"},{min:3,max:20,message:"长度在 3 到 10 个字符",trigger:"blur"}],password:[{required:!0,message:"请输入登陆密码",trigger:"blur"},{min:6,max:20,message:"长度在 6 到 15 个字符",trigger:"blur"}]};let x=I(),m=F();const y=async()=>{await u.value.validate();let l=await m.userLogin(t),o=await m.userInfo();l=="ok"?o=="ok"?(i({type:"success",message:"欢迎回来",title:`Hi, ${p()}好`}),x.push("/")):i({type:"error",message:o,title:`Hi, ${p()}好`}):i({type:"error",message:l,title:`Hi, ${p()}好`})};return(l,o)=>{const _=U,d=$,c=R,E=T,V=j,h=B;return C(),S("div",D,[e(h,null,{default:a(()=>[e(_,{span:16}),e(_,{span:8},{default:a(()=>[e(V,{class:"login_form",model:s(t),rules:w,ref_key:"loginForms",ref:u},{default:a(()=>[G,J,e(c,{prop:"username"},{default:a(()=>[e(d,{"prefix-icon":s(N),modelValue:s(t).username,"onUpdate:modelValue":o[0]||(o[0]=n=>s(t).username=n)},null,8,["prefix-icon","modelValue"])]),_:1}),e(c,{prop:"password"},{default:a(()=>[e(d,{type:"password","prefix-icon":s(q),modelValue:s(t).password,"onUpdate:modelValue":o[1]||(o[1]=n=>s(t).password=n),"show-password":""},null,8,["prefix-icon","modelValue"])]),_:1}),e(E,{class:"login_btn",type:"primary",size:"default",onClick:y},{default:a(()=>[H("登录")]),_:1})]),_:1},8,["model"])]),_:1})]),_:1})])}}});const Z=A(K,[["__scopeId","data-v-4a76de16"]]);export{Z as default};
