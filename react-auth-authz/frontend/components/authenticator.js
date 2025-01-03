import loadingStatus from "@/helpers/loadingStatus";
import { useUser } from "@/hooks/useUser";
import React from "react";
import { LoadingIndicator } from "./loadingIndicator";

export const Authenticator = () => {
  const { isAuthenticated, login, logout, getNameClaim, loadingState } =
    useUser();
  if (loadingState === loadingStatus.isLoading) {
    return <h4>Loading...</h4>;
  }
  if (isAuthenticated) {
    var username = getNameClaim();
    return (
      <div>
        Hi {username}
        <div>
          <button onClick={logout} className="mt-3 btn btn-secondary btn-sm">
            Logout
          </button>
        </div>
      </div>
    );
  } else {
    <button onClick={login} className="btn btn-primary">
      Login
    </button>;
  }
};
