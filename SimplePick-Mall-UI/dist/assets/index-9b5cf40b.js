import{at as e}from"./index-bc94ad6c.js";const g=(s,o,t)=>e.get("/api/sys/loginLog/list",{params:{current:s,pageSize:o,status:t}}),i=s=>e.post("/api/sys/loginLog/delete",s),r=(s,o,t)=>e.get("/api/sys/systemLog/list",{params:{current:s,pageSize:o,method:t}}),m=s=>e.post("/api/sys/systemLog/delete",s);export{i as a,r as b,m as c,g as r};