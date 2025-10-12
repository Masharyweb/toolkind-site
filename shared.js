const $=(s,r=document)=>r.querySelector(s); const $$=(s,r=document)=>Array.from(r.querySelectorAll(s));
let LANG=localStorage.getItem('ngokit_lang')||'ar';
function setLang(l){ LANG=l; localStorage.setItem('ngokit_lang',l);
  $$('[data-i18n]').forEach(el=>{ const k=el.getAttribute('data-i18n'); const txt=(I18N[k]&&I18N[k][LANG])||(I18N[k]&&I18N[k].ar)||'';
    if((el.tagName==='INPUT'||el.tagName==='TEXTAREA') && el.placeholder!==undefined){ el.placeholder=txt; } else { el.innerText=txt; }
  });
  document.documentElement.dir=(LANG==='ar'?'rtl':'ltr'); document.documentElement.lang=LANG;
}
document.addEventListener('DOMContentLoaded',()=>{ $('.lang-ar')?.addEventListener('click',()=>setLang('ar')); $('.lang-en')?.addEventListener('click',()=>setLang('en')); setLang(LANG); });
async function fileToImage(file){ return new Promise((resolve,reject)=>{ const url=URL.createObjectURL(file); const img=new Image(); img.onload=()=>{URL.revokeObjectURL(url); resolve(img)}; img.onerror=reject; img.src=url; }); }
