import React from 'react'
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'

import Home from './Home'
import NotFound from './NotFound'

class AppRouter extends React.Component {
  render() {
    return (
      <Router>
        <Switch>
          <Route exact path='/' component={Home} />
          <Route exact component={NotFound} />
        </Switch>
      </Router>
    )
  }
}

export default AppRouter
