function setActiveLink(event) {
  document.querySelectorAll(".filter-link").forEach((link) => {
    link.classList.remove("active");
  });
  event.currentTarget.classList.add("active");
}

function toggleSearchForm() {
  const searchForm = document.getElementById("searchForm");
  searchForm.style.display =
    searchForm.style.display === "block" ? "none" : "block";
}
