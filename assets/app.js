var vm = new Vue({
    el: "#app",
    data: {
        search: "",
        results: [],
    },
    delimiters: ['${', '}'],
    methods: {
        find: function(event) {
            // Check if valid siren or siret before actually making the request
            if (this.search.length < 9) {
                document.querySelector('.mdl-js-snackbar').MaterialSnackbar.showSnackbar({message: "Invalid siren : A siren must be 9 digits"}); 
                return
            }
            if ((this.search.length > 9 && this.search.length < 14) || this.search.length > 14) { 
                document.querySelector('.mdl-js-snackbar').MaterialSnackbar.showSnackbar({message: "Invalid siret : A siret must be 14 digits"}); 
                return
            }
            var siret = this.search.length > 9;
            var ep = this.$resource('/api/v1/'+(siret?'siret':'siren')+'/{id}');
            ep.get({id: this.search, limit: 100}).then(function(response) {
                this.results = siret?[response.data]:response.data;
            }, function(response) {
                console.log(response);
                document.querySelector('.mdl-js-snackbar').MaterialSnackbar.showSnackbar({message: "This "+ (siret?"siret":"siren") + " doesn't exist"});
            });
        }
    }
});