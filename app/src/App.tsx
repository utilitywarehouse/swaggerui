import { useEffect, useRef, useState } from "react";
import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

function App() {
  // This nonsense is necessary due to a closure in the requestInterceptor
  // which will always use the default `token` value if trying to just use
  // state.
  const [token, setToken] = useState("");
  const tokenRef = useRef("");
  useEffect(() => {
    tokenRef.current = token;
  }, [token]);

  return (
    <div>
      {/* Very barebones, but it works :shrug: */}
      <h2>Okta token:</h2>
      <input
        type="text"
        value={token}
        onChange={(e) => {
          setToken(e.target.value);
        }}
      />
      <SwaggerUI
        requestInterceptor={(req) => ({
          ...req,
          headers: {
            Authorization: `Bearer ${tokenRef.current.trim()}`,
          },
        })}
        url={`${window.location.origin}/swagger.json`}
      />
    </div>
  );
}

export default App;
