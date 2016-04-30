describe('SearchController tests', function() { 	
  test_init()

  // mock service
  beforeEach(mockState);

  var searchService = {
      Search: function (query) {
        deferred = q.defer();
        return deferred.promise;
    },
  }

  searchCtl = {};

  beforeEach(inject(function($controller, $rootScope, $q, $route, $state){
    q = $q;
    realState = $state
    route = $route
    scope = $rootScope.$new();
    searchCtl = $controller('searchCtl', {
      $scope: scope,
      searchService: searchService,
      $state: state
    })
  }))

  it('default action should execute a search without query or params', function(){
    spyOn(searchService, 'Search').and.callThrough();
    scope.init();
    expect(searchService.Search).toHaveBeenCalledWith({})
 });

  it('default action should call a method to process searchResults', function(){
    scope.init();
    deferred.resolve('ok');
    scope.$root.$digest();
    expect(scope.data).toBe('ok');
  });

  it('should display error message if default action fails', function(){
    scope.init();
    deferred.reject('error');
    scope.$root.$digest();
    expect(scope.msg.error).toBe('error');
  });

  it('should contain a search method', function(){
    expect(scope.search).not.toBe(undefined);
  });

  it('search method should call the search service', function(){
    spyOn(searchService, 'Search').and.callThrough();
    scope.search();
    expect(searchService.Search).toHaveBeenCalled()
  });

  it('search method should get the query from scope', function(){
    spyOn(searchService, 'Search').and.callThrough();
    scope.query = 'abstract'
    scope.search();
    expect(searchService.Search).toHaveBeenCalledWith({"query":"abstract"})
  });

  it('search method should get page from scope', function(){
    spyOn(searchService, 'Search').and.callThrough();
    scope.page = 2
    scope.search();
    expect(searchService.Search).toHaveBeenCalledWith({"page":2})
  });

  it('search method allow filter by facet as param', function(){
    spyOn(searchService, 'Search').and.callThrough();
    scope.search({'category': ['2']});
    expect(searchService.Search).toHaveBeenCalledWith({"filters":{'category': ['2']}})
  });

  it('should generate pagination info', function(){
    scope.init();
    result = {"navigation":{"order":"","currentPage":1,"pageSize":0,"pageStart":1,"pageEnd":20,
      "totalPages":136,"totalCount":2713}};

    deferred.resolve(result);
    scope.$root.$digest();

    p = scope.pagination;
    expect(p.page).toBe(1);
    expect(p.total_pages).toBe(136);
    expect(p.total_results).toBe(2713);
    expect(p.start).toBe(1);
    expect(p.end).toBe(20);
    expect(p.min_page).toBe(1);
    expect(p.max_page).toBe(15);
    expect(p.pages.length).toBe(15);
  });
});
