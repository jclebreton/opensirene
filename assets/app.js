// window.addEventListener("load", function(event) {

var mdlLoaded = setInterval(function(){ 
    if(typeof document.querySelector('.mdl-js-layout') !== 'undefined') init()
}, 100)

function init () { 
    clearInterval(mdlLoaded)
}
    var vm = new Vue({
        el: "#app",
        data: {
            search: "",
            results: [],
        },
        delimiters: ['${', '}'],
        methods: {
            find: function(event) {
                if (this.search.length < 8) { 
                    document.querySelector('.mdl-js-snackbar').MaterialSnackbar.showSnackbar({message: "Invalid siren or siret"}); 
                    return
                }
                var siret = this.search.length > 9;
                var ep = this.$resource('/api/v1/siren/{id}');
                if (siret) {
                    ep = this.$resource('/api/v1/siret/{id}');
                }
                ep.get({id: this.search, limit: 100}).then(function(response) {
                    this.results = siret?[response.data]:response.data;
                }, function(response) {
                    document.querySelector('.mdl-js-snackbar').MaterialSnackbar.showSnackbar({message: "This "+ (siret?"siret":"siren") + " doesn't exist"});
                });
            }
        }
    });
// });