const Login = () => {
    const handleLogin = () => {
        // Start the login flow by redirecting to the backend
        window.location.href = 'http://localhost:8080/login/google'; // Use a GET request for redirection
      };
    
      return (
        <button onClick={handleLogin}>
          Login with Google
        </button>
      );
};

export default Login;