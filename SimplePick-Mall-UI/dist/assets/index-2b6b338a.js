import{at as t}from"./index-bc94ad6c.js";const r=()=>t.post("/api/pms/category/list"),p=e=>e.id?t.post("/api/pms/category/update",e):t.post("/api/pms/category/add",e),o=e=>t.post("/api/pms/category/delete",e);export{p as a,o as b,r};