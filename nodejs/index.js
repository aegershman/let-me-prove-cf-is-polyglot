const app = require('express')()

app.get('/', (req, res) => {
    res.send(`
        Hey there! I'm a nodejs app.
        Look at me go.
        `)
})

app.listen(process.env.PORT || 3000);
