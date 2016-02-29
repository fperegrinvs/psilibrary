var Psilibrary;
(function (Psilibrary) {
    angular.module('psilibrary.config')
        .constant('config', {
        'serverUrl': 'http://localhost:8080',
        'retries': 3
    });
})(Psilibrary || (Psilibrary = {}));
