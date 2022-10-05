import React, { useCallback, useEffect, useState } from "react";
import logo from "./eli-logo.svg";
import "./App.css";
import axios from "axios";

function App() {
  const [systems, setSystems] = useState([]);
  const [searchText, setSearchText] = useState("");


  const getData = useCallback(
    () => {
      let url = "http://localhost:3700/v1/systems?limit=50";
      if (searchText !== "") url += "&searchText=" + searchText;
      axios
        .get(url)
        .then((res) => {
          setSystems(res.data);
        })
        .catch((err) => {
          console.error(err);
        })
        .finally
        //could be end of busy indicator here
        ();
    },
    [searchText],
  );

  useEffect(() => {
    getData()
    document.title = "Systems database";
  });



  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
      </header>
      <body>
        <div className="systemsHeader">Systems</div>
        <div className="systemsSearch">
          <input
            type="text"
            value={searchText}
            placeholder="Search by name or code"
            onChange={(e) => {
              setSearchText(e.target.value);
            }}
          />
        </div>
        <div className="systemsWrapper">
          {systems.map((val: any) => (
            <div className="systemsItem" key={val.code}>
              <div>Name: {val.name}</div>
              <div>Code:{val.code}</div>
              <div>Parent:{val.parentSystemCode}</div>
            </div>
          ))}
        </div>
      </body>
    </div>
  );
}

export default App;
