import SwaggerUI from "swagger-ui-react";
import "swagger-ui-react/swagger-ui.css";

function App() {
  const OKTA_TOKEN_GOES_HERE = "super secret token";
  return (
    <div>
      <SwaggerUI
        requestInterceptor={(req) => ({
          ...req,
          headers: {
            Authorization: `Bearer ${OKTA_TOKEN_GOES_HERE}`,
          },
        })}
        url={`${window.location.origin}/swagger.json`}
      />
    </div>
  );
}

export default App;
