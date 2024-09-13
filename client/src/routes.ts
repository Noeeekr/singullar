const routes = {
    auth: {
        main: "/auth",
        recover: "/recover"
    },
    public: [
        '/dummyRoute'
    ],
    private: [

    ],
    // private: {} is not really necessary, since I have
    // few public routes I can just filter them and make 
    // the default for a route to be a private
}

export default routes