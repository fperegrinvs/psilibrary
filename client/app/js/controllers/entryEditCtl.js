'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryEditCtl', ['$scope', 'entryService', '$state', 'categoryService', 
    	function ($scope, entryService, $state, categoryService) {

	$scope.init = function(){
        if (!$state.params.id){
            $state.go('entryList', {error: 'Registro não encontrado'});
        }

        entryService.Get($state.params.id);
        $scope.categories = categoryService.List();
	}

	$scope.save = function(){
		entryService.Update($scope.data).then(
			function(data){
				$scope.msg = {success: 'Registro atualizado com sucesso'};
			},
			function(err){
				$scope.msg = {error: err};
			});
	}

	$scope.init();
}]);