'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryCreateCtl', ['$scope', 'entryService', '$state', 'categoryService', 
        function ($scope, entryService, $state, categoryService) {    	
    	$scope.init = function(){
            $scope.categories = categoryService.List();
    	}

    	$scope.save = function(){
            $scope.$broadcast('show-errors-check-validity');
            if ($scope.dataForm && !$scope.dataForm.$valid) {
                return
            }

    		entryService.Create($scope.data).then(
    			function(data){
    				$scope.msg = {success: "Registro inserido com sucesso"};
    			}, 
	    		function(err){
	    			$scope.msg = {error: err};
	    		});
    	}

      $scope.init();
}]);