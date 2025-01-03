import currencyFormatter from "@/helpers/currencyFormatter";
import React, { useContext } from "react";
import { navigationContext } from "./app";
import navValues from "@/helpers/navValues";

const HouseRow = ({ house }) => {
  const { navigate } = useContext(navigationContext);
  return (
    <tr onClick={() => navigate(navValues.house, house)}>
      <td>{house.address}</td>
      <td>{house.country}</td>
      <td>{currencyFormatter.format(house.price)}</td>
    </tr>
  );
};

const HouseRowMem = React.memo(HouseRow);
export default HouseRow;
export { HouseRowMem };
