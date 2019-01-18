import React from 'react'
import Loadable from 'react-loadable'

import Loading from '../../components/loading'

const LoadableComponent = Loadable({
  loader: () => import('./Home'),
  loading: Loading
})

const LoadableHome = () => {
  return <LoadableComponent />
}

export default LoadableHome
