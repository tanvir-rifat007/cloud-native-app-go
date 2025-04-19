export const API = {
  baseURL: "http://localhost:8080/api/v1",

  postNewsletter: async (body) => {
    return API.fetchData("/newsletter", body);
  },

  fetchData: async (url, data = {}) => {
    try {
      const response = await fetch(API.baseURL + url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });

      if (!response.ok) {
        return {
          error: {
            status: response.status,
            message: response.statusText,
          },
        };
      }
      const res = await response.json();

      return res;
    } catch (error) {
      console.error("Error fetching data:", error);
      throw error;
    }
  },
};
