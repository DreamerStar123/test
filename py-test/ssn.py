import pandas as pd
import json

# Specify the path to your Excel file
file_path = 'E:/doc/Texas SSNs.xlsx'

# Read the Excel file
df = pd.read_excel(file_path, sheet_name='Sheet1')  # You can specify the sheet name or index

# Display the first few rows of the DataFrame
n = 104
print('{} rows'.format(len(df)))

with open('row/row_data{}.json'.format(n), 'w') as json_file:
    json.dump(json.loads(df.iloc[n].to_json()), json_file, indent=4)

# for n in range(0, 200):
#     # Save the JSON object to a file
#     with open('row/row_data{}.json'.format(n), 'w') as json_file:
#         json.dump(json.loads(df.iloc[n].to_json()), json_file, indent=4)
