var Psilibrary;
(function (Psilibrary) {
    angular.module('psilibrary.config')
        .constant('config', {
        'serverUrl': 'http://52.67.51.36',
//        'serverUrl': 'http://localhost:8080',
        'retries': 3
    });
})(Psilibrary || (Psilibrary = {}));
