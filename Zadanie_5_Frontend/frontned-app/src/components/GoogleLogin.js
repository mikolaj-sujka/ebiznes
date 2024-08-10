import React, { useEffect } from "react";
import { useAuth } from "../AuthContext";

const GoogleLogin = () => {
  const { googleLogin } = useAuth();

  const handleLoginClick = () => {
    window.location.href = "http://localhost:8080/login/google";
  };

  useEffect(() => {
    const fetchToken = async () => {
      const params = new URLSearchParams(window.location.search);
      const code = params.get("code");
      const state = params.get("state");

      if (code && state === "state-token") {
        try {
          const response = await fetch(
            "http://localhost:8080/auth/callback?code=" + code + "&state=" + state,
            {
              method: "GET",
              credentials: "include",
            }
          );
          const data = await response.json();

          if (response.ok) {
            // Save token to local storage
            localStorage.setItem("token", data.token);
            googleLogin(data.token);
          } else {
            console.error("Login failed:", data.error);
          }
        } catch (error) {
          console.error("Error during authentication:", error);
        }
      }
    };

    fetchToken();
  }, [googleLogin]);

  return (
    <div>
      <button onClick={handleLoginClick}>Login with Google</button>
    </div>
  );
};

export default GoogleLogin;
