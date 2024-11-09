function setActiveLink(event) {
  document.querySelectorAll(".filter-link").forEach((link) => {
    link.classList.remove("active");
  });
  event.currentTarget.classList.add("active");
}
