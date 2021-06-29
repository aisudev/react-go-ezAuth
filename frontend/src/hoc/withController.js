const withController = (Component) => (Controller) => () => {
    return (
        <Controller>
            <Component />
        </Controller>
    )
}

export default withController
