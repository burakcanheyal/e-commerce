import React, { Suspense } from "react";

const Loadable = (Component) => (props) => (
    <Suspense fallback={null}>
        <Component {...props} />
    </Suspense>
);

export default Loadable;
