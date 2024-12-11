import defaultPhoto from "@/helpers/defaultPhoto";
import React from "react";

export const House = ({ house }) => {
  return (
    <div className="row">
      <div className="col-6">
        <div className="row">
          <img
            className="img-fluid"
            src={
              house.photo ? `./houseImages/${house.photo}.jpeg` : defaultPhoto
            }
            alt="House pic"
          />
        </div>
      </div>
    </div>
  );
};
