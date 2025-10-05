// ToolKind dynamic script

// 🔄 تبديل اللغة
const langBtn = document.getElementById("langToggle");
langBtn.addEventListener("click", () => {
  const currentLang = document.documentElement.lang;
  if (currentLang === "ar") {
    document.documentElement.lang = "en";
    document.querySelector("h2").textContent = "Welcome to ToolKind";
    document.querySelector("p").textContent =
      "A smart tool library to help you complete your daily tasks efficiently.";
    document.querySelector(".btn").textContent = "Explore Tools";
    langBtn.textContent = "العربية 🇸🇦";
  } else {
    document.documentElement.lang = "ar";
    document.querySelector("h2").textContent = "مرحبًا بك في ToolKind";
    document.querySelector("p").textContent =
      "مكتبة أدوات ذكية تساعدك في إنجاز مهامك اليومية بسرعة وأناقة.";
    document.querySelector(".btn").textContent = "استكشف الأدوات";
    langBtn.textContent = "English 🇬🇧";
  }
});

// 🌙 الوضع الليلي التلقائي
const prefersDark = window.matchMedia("(prefers-color-scheme: dark)");
if (prefersDark.matches) {
  document.body.classList.add("dark");
}

prefersDark.addEventListener("change", (e) => {
  if (e.matches) {
    document.body.classList.add("dark");
  } else {
    document.body.classList.remove("dark");
  }
});
