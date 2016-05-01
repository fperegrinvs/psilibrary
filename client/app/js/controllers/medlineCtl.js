'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('medlineCtl', ['$scope', 'medlineService', '$state', function ($scope, medlineService, $state) {
        window.currentScope = $scope;
        $scope.import = function () {
            medlineService.Import($scope.xml).then(function () {
            });
        };
    }]);
