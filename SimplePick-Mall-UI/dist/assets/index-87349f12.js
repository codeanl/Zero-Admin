import{at as e}from"./index-bc94ad6c.js";const o=(s,t,r,p)=>e.get("/api/sms/subject/list",{params:{current:s,pageSize:t,name:r,status:p}}),u=s=>e.post("/api/sms/subject/delete",s),c=s=>s.id?e.post("/api/sms/subject/update",s):e.post("/api/sms/subject/add",s);export{u as a,c as b,o as r};