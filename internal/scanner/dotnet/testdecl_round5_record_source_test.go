//ff:type feature=scan type=test topic=dotnet
//ff:what round5RecordSource 테스트 보조 선언
package dotnet

// Controller exercising: records, data annotations, multiple attributes,
// array/nullable/generic property types, authorize roles.
var round5RecordSource = `
namespace MyApp.Models;

public record CreateItemRequest(
    [Required] string Name,
    int? Quantity,
    List<string> Tags,
    decimal[] Prices);

public class ItemResponse
{
    public string Id { get; set; }
    public DateTime? CreatedAt { get; set; }
    public List<TagDto> Tags { get; set; }
}

public class TagDto
{
    public string Label { get; set; }
}
`
