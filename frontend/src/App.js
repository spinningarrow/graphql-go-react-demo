import React, { Component } from "react";
import logo from "./logo.svg";
import "./App.css";
import { graphql, QueryRenderer } from "react-relay";
import environment from "./relay";

class App extends Component {
  render() {
    return (
      <QueryRenderer
        environment={environment}
        query={graphql`
          query AppQuery {
            hello
          }
        `}
        variables={{}}
        render={({ error, props }) => {
          if (error) {
            return <div>ERRORRR</div>;
          }
          if (!props) {
            return <div>Loading...</div>;
          }
          return (
            <div className="App">
              <header className="App-header">
                <img src={logo} className="App-logo" alt="logo" />
                <p>{props.hello}</p>
                <p>{props.anotherOne ? "YES" : "NO"}</p>
                <a
                  className="App-link"
                  href="https://reactjs.org"
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Learn React
                </a>
              </header>
            </div>
          );
        }}
      />
    );
  }
}

export default App;
