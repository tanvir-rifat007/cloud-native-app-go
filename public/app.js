import { API } from "./services/API.js";
import Router from "./services/Router.js";

globalThis.addEventListener("DOMContentLoaded", () => {
  app.router.init();
});

globalThis.app = {
  signupNewsletter: async (event) => {
    const email = document.querySelector("#email").value;

    event.preventDefault();

    const response = await API.postNewsletter({
      email,
    });

    console.log(response);
  },
  router: Router,
};
