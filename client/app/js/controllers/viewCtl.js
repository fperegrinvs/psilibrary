'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('viewCtl', ['$scope', 'entryService', '$state', 'categoryService', 'entryTypeService', '$sce', 
    	function ($scope, entryService, $state, categoryService, entryTypeService, $sce) {

	$scope.init = function(){
        $scope.section = 'entry';

        if (!$state.params.id){
            $state.go('entry', {error: 'Registro não encontrado'});
        }

        var entryCall = entryService.Get($state.params.id);
        entryCall.then(function(data){
            $scope.data = data;
            $scope.data.publishDateLocal = new Date(data.publishDate);
            $scope.data.abstract_safe = $sce.trustAsHtml(data.abstract);
            $scope.data.content_safe = $sce.trustAsHtml(data.content);
            $scope.current = data.title;
        },
        function(err){
            $scope.msg = {error: err}
        });
	}

	$scope.init();
}]);