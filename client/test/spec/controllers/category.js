describe('category tests', function() {
  beforeEach(module('client'))
  describe('list', function(){
  	var listCtl, scope, list, service, q, deferred, state, currentState

  	beforeEach(function(){
  		service = {
            List: function () {
            	deferred = q.defer();
                return deferred.promise;
            },
            Get: function (id) {
            	deferred = q.defer();
                return deferred.promise;
            },
            Update: function (category) {
            	deferred = q.defer();
                return deferred.promise;
            },
            Create: function (category) {
            	deferred = q.defer();
                return deferred.promise;
            }

  		}

  		state = {
  			Go: function(s) {
  				currentState = s;
  			}
  		}
  	});


  	beforeEach(inject(function($controller, $rootScope, $q){
  		q = $q;
  		scope = $rootScope.$new();
  		listCtl = $controller('categoryListCtl', {
  			$scope: scope,
  			categoryService: service,
  			$state: state
  		})
  	}))

  
	it('should get list during init', function() {  	
		spyOn(service, 'List').and.callThrough();

		scope.init();

		deferred.resolve({});
		scope.$root.$digest();

		expect(service.List).toHaveBeenCalled();
	});

  })
});