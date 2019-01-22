import * as React from 'react'
import { Dispatch } from 'redux'
import { connect } from 'react-redux'

import { map } from 'lodash'

import { RootState, projects } from '../redux'

type Props = {
    dispatch: Dispatch
}

type ReduxProps = {
    entries: projects.ProjectsList
}

class ProjectsList extends React.Component<Props & ReduxProps> {
    componentDidMount() {
        const { dispatch } = this.props
        dispatch(projects.list())
    }

    render() {
        const { entries } = this.props
        return (
            <div>
                <h4>Projects</h4>
                <ul>
                    { map(entries, (entry) => (
                        <li key={ entry.getName() }>{ entry.getName() }</li>
                    )) }
                </ul>
            </div>
        )
    }
}

const mapStateToProps = (state: RootState): ReduxProps => {
    return {
        entries: projects.getAll(state)
    }
}

export default connect(mapStateToProps)(ProjectsList)