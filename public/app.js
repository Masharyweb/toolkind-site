document.addEventListener('click', (e)=>{
  const sb = document.querySelector('.sidebar');
  if(!sb) return;
  if(e.target.closest('[data-menu]')) sb.classList.add('open');
  if(e.target.closest('[data-close]') || e.target === sb) sb.classList.remove('open');
});
const html=document.documentElement;const b=document.querySelector('[data-lang]');if(b){b.addEventListener('click',()=>{const n=html.lang==='ar'?'en':'ar';html.lang=n;html.dir=n==='ar'?'rtl':'ltr';localStorage.setItem('lng',n);document.querySelectorAll('[data-t]').forEach(el=>{const k=el.getAttribute('data-t');el.textContent=(window.i18n[n]&&window.i18n[n][k])||el.textContent;});});const s=localStorage.getItem('lng')||'ar';html.lang=s;html.dir=s==='ar'?'rtl':'ltr';}
